package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/sunyatsuntobee/server/models"
)

func initCollectionActivityStagesRouter(router *mux.Router) {
	// GET /activity_stages
	router.HandleFunc("/api/activity_stages", activityStagesGetHandler()).
		Methods(http.MethodGet)

	// GET /activity_stages/yy-mm-dd
	router.HandleFunc("/api/activity_stages/{date}",
		activityStagesDateGetHandler()).Methods(http.MethodGet)
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
