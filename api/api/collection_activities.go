package api

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/sunyatsuntobee/server/models"
	"github.com/sunyatsuntobee/server/util"
)

func initCollectionActivitiesRouter(router *mux.Router) {
	// GET /activities
	router.HandleFunc("/api/activities", activitiesGetHandler()).
		Methods(http.MethodGet)
	// DELETE /activities/{ID}/stages
	router.HandleFunc("/api/activities/{ID}/stages",
		activityStagesDeleteHandler()).Methods(http.MethodDelete)
	// POST /activities/{ID}/stages
	router.HandleFunc("/api/activities/{ID}/stages",
		activityStagesPostHandler()).Methods(http.MethodPost)
	// PUT /activities
	router.HandleFunc("/api/activities/{ID}", activitiesPutHandler()).
		Methods(http.MethodPut)
}

func activityStagesPostHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		id, _ := strconv.Atoi(mux.Vars(req)["ID"])
		stageNum, _ := strconv.Atoi(req.FormValue("stage_num"))
		startTime, _ := time.Parse("2006-01-02T15:04:05", req.FormValue("start_time"))
		endTime, _ := time.Parse("2006-01-02T15:04:05", req.FormValue("end_time"))
		stage := models.ActivityStage{
			StageNum:   stageNum,
			StartTime:  startTime,
			EndTime:    endTime,
			Location:   req.FormValue("location"),
			Content:    req.FormValue("content"),
			ActivityID: id,
		}
		models.ActivityStageDAO.InsertOne(&stage)
		formatter.JSON(w, http.StatusCreated, nil)
	}

}

func activityStagesDeleteHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(req)["ID"])
		models.ActivityStageDAO.DeleteByAID(id)
		formatter.JSON(w, http.StatusNoContent, nil)
	}

}

func activitiesPutHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		id, _ := strconv.Atoi(mux.Vars(req)["ID"])
		activity, _ := models.ActivityDAO.FindByID(id)
		activity.Name = req.FormValue("name")
		activity.Description = req.FormValue("description")
		activity.Category = req.FormValue("category")
		logoPath := "static/assets/activities/logos/" + strconv.Itoa(id) +
			".png"
		if !strings.HasPrefix(req.FormValue("logo_url"), "/") {
			util.SaveBase64AsPNG(req.FormValue("logo_url"), logoPath)
		}
		posterPath := "static/assets/activities/posters/" + strconv.Itoa(id) +
			".png"
		if !strings.HasPrefix(req.FormValue("poster_url"), "/") {
			util.SaveBase64AsPNG(req.FormValue("poster_url"), posterPath)
		}

		activity.LogoURL = "/" + logoPath
		activity.PosterURL = "/" + posterPath
		activity.OrganizationID, _ = strconv.Atoi(
			req.FormValue("organization_id"))
		models.ActivityDAO.UpdateOne(&activity)
		formatter.JSON(w, http.StatusCreated, nil)
	}

}

func activitiesGetHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		w.Header().Set("Access-Control-Allow-Origin", cor)
		if req.FormValue("oid") == "" {

		} else {
			oid, _ := strconv.Atoi(req.FormValue("oid"))
			data := models.ActivityDAO.FindByOID(oid)
			formatter.JSON(w, http.StatusOK, data)
		}

	}

}
