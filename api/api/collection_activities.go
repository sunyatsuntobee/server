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

	// POST /activities/{ID}/stages
	router.HandleFunc(url+"/{ID}/stages",
		activityStagesPostHandler()).Methods(http.MethodPost)

	// PUT /activities/{ID}
	router.HandleFunc(url+"/{ID}", activitiesPutHandler()).
		Methods(http.MethodPut)
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
		formatter.JSON(w, http.StatusCreated, nil)
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

	}

}
