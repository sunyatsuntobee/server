package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/sunyatsuntobee/server/logger"
	"github.com/sunyatsuntobee/server/models"
)

func initCollectionActivityStagesRouter(router *mux.Router) {
	url := "/api/activity_stages"

	// GET /activity_stages
	router.HandleFunc(url, activityStagesGetHandler()).
		Methods(http.MethodGet)

	// GET /activity_stages/yy-mm-dd
	router.HandleFunc(url+"/{date}",
		activityStagesDateGetHandler()).Methods(http.MethodGet)
	
	// PUT /activity_stages/{ID}
	router.HandleFunc(url, activityStagesPutHandler()).
		Methods(http.MethodPut)
}

func activityStagesDateGetHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		date, err := time.Parse(dateLayout, mux.Vars(req)["date"])
		if err != nil {
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("Bad Request", "日期格式错误", nil))
			return
		}
		data := models.ActivityStageDAO.FindFullByDay(date)
		formatter.JSON(w, http.StatusOK,
			NewJSON("OK", "成功获取活动列表", data))
	}

}

func activityStagesGetHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		valAid := req.FormValue("aid")
		valDate := req.FormValue("one_day_in_that_month")

		if valAid != "" {
			aid, _ := strconv.Atoi(valAid)
			data := models.ActivityStageDAO.FindByAID(aid)
			formatter.JSON(w, http.StatusOK,
				NewJSON("OK", "成功获取活动列表", data))
		} else if valDate != "" {
			date, err := time.Parse("06-01-02", valDate)
			if err != nil {
				formatter.JSON(w, http.StatusBadRequest,
					NewJSON("Bad Request", "日期格式错误", nil))
				return
			}
			data := models.ActivityStageDAO.FindFullByMonth(date)
			formatter.JSON(w, http.StatusOK,
				NewJSON("OK", "成功获取活动列表", data))
		}
	}

}

func activityStagesPutHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		var t models.ActivityStage
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&t)
		if err != nil {
			logger.E.Println(err)
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "数据格式错误", nil))
			return
		}
		_, has := models.ActivityStageDAO.FindByID(t.ID)
		if (!has) {
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("created", "无此活动阶段", nil))
		} else {
			models.ActivityStageDAO.UpdateOne(&t)
			formatter.JSON(w, http.StatusOK,
				NewJSON("ok", "成功修改活动阶段", t))
		}

	}

}

