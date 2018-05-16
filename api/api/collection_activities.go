package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sunyatsuntobee/server/logger"
	"github.com/sunyatsuntobee/server/models"
)

func initCollectionActivitiesRouter(router *mux.Router) {
	url := "/api/activities"

	// GET /activities{?oid}
	router.HandleFunc(url,
		activitiesGetHandler()).Methods(http.MethodGet)

	// GET /activities{?actid}
	router.HandleFunc(url,
		activitiesGetHandler()).Methods(http.MethodGet)

	// POST /activities/{ID}/stages
	router.HandleFunc(url+"/{ID}/activity_stages",
		activityStagesPostHandler()).Methods(http.MethodPost)

	// POST /activities
	router.HandleFunc(url,
		activitiesCreateHandler()).Methods(http.MethodPost)

	// PUT /activities/{ID}
	router.HandleFunc(url+"/{ID}", activitiesPutHandler()).
		Methods(http.MethodPut)
}

func activitiesCreateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		decoder := json.NewDecoder(req.Body)
		var activity models.Activity
		err := decoder.Decode(&activity)
		if err != nil {
			logger.E.Println(err)
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "数据格式错误", nil))
			return
		}

		models.ActivityDAO.InsertOne(&activity)
		formatter.JSON(w, http.StatusCreated,
			NewJSON("Created", "创建新的活动成功", activity))
	}
}

func activityStagesPostHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		var stage models.ActivityStage
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&stage)
		if err != nil {
			logger.E.Println(err)
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "数据格式错误", nil))
			return
		}
		activityID, _ := strconv.Atoi(mux.Vars(req)["ID"])
		stage.ActivityID = activityID
		models.ActivityStageDAO.InsertOne(&stage)
		formatter.JSON(w, http.StatusOK,
			NewJSON("created", "活动阶段创建成功", stage))
	}

}

func activitiesPutHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		var activity models.Activity
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&activity)
		if err != nil {
			logger.E.Println(err)
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "数据格式错误", nil))
			return
		}
		id, _ := strconv.Atoi(mux.Vars(req)["ID"])
		old, _ := models.ActivityDAO.FindByID(id)
		activity.LogoURL = old.LogoURL
		activity.PosterURL = old.PosterURL
		models.ActivityDAO.UpdateOne(&activity)
		formatter.JSON(w, http.StatusCreated,
			NewJSON("created", "修改活动信息成功", activity))
	}

}

func activitiesGetHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if req.FormValue("oid") == "" {

		} else {
			oid, _ := strconv.Atoi(req.FormValue("oid"))
			data := models.ActivityDAO.FindByOID(oid)
		
			
			formatter.JSON(w, http.StatusOK,
				NewJSON("OK", "获取活动列表成功", data))
		}

		if req.FormValue("actid") == "" {

			} else {
				actid, _ := strconv.Atoi(req.FormValue("actid"))
				data,_:= models.ActivityDAO.FindFullByAID(actid)
				if data.ID == 0 {
					data.Name = "该活动尚未录入信息"
					data.ShortName = "该活动尚未录入信息"
					data.Description = "可以到图蜂后台录入信息"
				}
				formatter.JSON(w, http.StatusOK,
					NewJSON("OK", "获取活动信息成功", data))
			}
	}

}
