package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sunyatsuntobee/server/models"
	"github.com/sunyatsuntobee/server/util"
)

func initCollectionPhotoLivesRouter(router *mux.Router) {
	url := "/api/photolives"
	router.HandleFunc(url, photoLivesPostHandler()).Methods(http.MethodPost)
	router.HandleFunc(url, photoLivesPutHandler()).Methods(http.MethodPut)
	router.HandleFunc(url, optionsHandler()).Methods(http.MethodOptions)
}

func photoLivesPostHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		w.Header().Set("Access-Control-Allow-Origin", cor)
		photoLive, err := obtainPhotoLiveFromRequest(req)
		if err != nil {
			formatter.JSON(w, http.StatusBadRequest, util.Error{
				Msg: err.Error(),
			})
			return
		}
		models.PhotoLiveDAO.InsertOne(&photoLive)
		supervisors := strings.Split(req.FormValue("supervisor_ids"), ",")
		for i := range supervisors {
			supervisorID, _ := strconv.Atoi(supervisors[i])
			if supervisorID == 0 {
				continue
			}
			relationship := models.PhotoLivesSupervisors{
				PhotoLiveID:  photoLive.ID,
				SupervisorID: supervisorID,
			}
			models.PhotoLivesSupervisorsDAO.InsertOne(&relationship)
		}
		formatter.JSON(w, http.StatusCreated, nil)
	}

}

func photoLivesPutHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		w.Header().Set("Access-Control-Allow-Origin", cor)
		photoLive, err := obtainPhotoLiveFromRequest(req)
		if err != nil {
			formatter.JSON(w, http.StatusBadRequest, util.Error{
				Msg: err.Error(),
			})
			return
		}
		photoLive.ID, _ = strconv.Atoi(req.FormValue("id"))
		models.PhotoLiveDAO.UpdateOne(&photoLive)
		supervisors := strings.Split(req.FormValue("supervisor_ids"), ",")
		models.PhotoLivesSupervisorsDAO.ClearByPLID(photoLive.ID)
		for i := range supervisors {
			supervisorID, _ := strconv.Atoi(supervisors[i])
			if supervisorID == 0 {
				continue
			}
			relationship := models.PhotoLivesSupervisors{
				PhotoLiveID:  photoLive.ID,
				SupervisorID: supervisorID,
			}
			models.PhotoLivesSupervisorsDAO.InsertOne(&relationship)
		}
		formatter.JSON(w, http.StatusCreated, nil)
	}

}

func obtainPhotoLiveFromRequest(req *http.Request) (models.PhotoLive, error) {
	expectMembers, _ := strconv.Atoi(req.FormValue("expect_members"))
	activityStageID, _ := strconv.Atoi(req.FormValue("activity_stage_id"))
	managerID, _ := strconv.Atoi(req.FormValue("manager_id"))
	photographerManagerID, _ := strconv.Atoi(
		req.FormValue("photographer_manager_id"))
	if activityStageID == 0 || managerID == 0 ||
		photographerManagerID == 0 {
		return models.PhotoLive{}, errors.New("请填写必需字段")
	}
	photoLive := models.PhotoLive{
		ExpectMembers:         expectMembers,
		AdProgress:            req.FormValue("ad_progress"),
		ActivityStageID:       activityStageID,
		ManagerID:             managerID,
		PhotographerManagerID: photographerManagerID,
	}
	return photoLive, nil
}
