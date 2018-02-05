package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/sunyatsuntobee/server/models"
)

func initCollectionPhotosRouter(router *mux.Router) {
	url := "/api/photos"
	router.HandleFunc(url, photosGetHandler()).Methods(http.MethodGet)
	router.HandleFunc(url, photosPutHandler()).Methods(http.MethodPut)
}

func photosGetHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if req.FormValue("category") == "" {
			data := models.PhotoDAO.FindFullAll()
			formatter.JSON(w, http.StatusOK, data)
		} else {
			data := models.PhotoDAO.FindFullByCategory(
				req.FormValue("category"))
			formatter.JSON(w, http.StatusOK, data)
		}
	}

}

func photosPutHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		id, _ := strconv.Atoi(req.FormValue("id"))
		old, _ := models.PhotoDAO.FindByID(id)
		old.URL = req.FormValue("url")
		old.TookTime, _ = time.Parse(time.RFC3339, req.FormValue("took_time"))
		old.TookLocation = req.FormValue("took_location")
		old.ReleaseTime, _ = time.Parse(time.RFC3339,
			req.FormValue("release_time"))
		old.Category = req.FormValue("category")
		old.Likes, _ = strconv.Atoi(req.FormValue("likes"))
		old.RejectReason = req.FormValue("reject_reason")
		old.PhotographerID, _ = strconv.Atoi(req.FormValue("photographer_id"))
		models.PhotoDAO.UpdateOne(&old)
		formatter.JSON(w, http.StatusCreated, nil)
	}

}
