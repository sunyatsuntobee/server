package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sunyatsuntobee/server/models"
)

func initConsolePhotoLivesRouter(router *mux.Router) {
	router.HandleFunc("/photolives", consolePhotoLivesGetHandler()).
		Methods(http.MethodGet)
	router.HandleFunc("/photolives/add", consolePhotoLivesAddGetHandler()).
		Methods(http.MethodGet)
	router.HandleFunc("/photoLives/add", consolePhotoLivesAddPostHandler()).
		Methods(http.MethodPost)
}

func consolePhotoLivesGetHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if req.FormValue("id") == "" {
			data := models.PhotoLiveDAO.FindFullAll()
			formatter.HTML(w, http.StatusOK,
				"console/photo_lives/photo_lives_list", data)
		} else {
			id, _ := strconv.Atoi(req.FormValue("id"))
			data := models.PhotoLiveDAO.FindFullByID(id)
			formatter.HTML(w, http.StatusOK,
				"console/photo_lives/photo_live", data)
		}
	}

}

func consolePhotoLivesAddGetHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		data := models.OrganizationDAO.FindAll()
		formatter.HTML(w, http.StatusOK,
			"console/photo_lives/photo_live_add", data)
	}

}

func consolePhotoLivesAddPostHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
	}

}
