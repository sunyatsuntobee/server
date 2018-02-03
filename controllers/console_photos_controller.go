package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sunyatsuntobee/server/models"
)

func initConsolePhotosRouter(router *mux.Router) {
	router.HandleFunc("/photos", consolePhotosGetHandler()).
		Methods(http.MethodGet)
}

func consolePhotosGetHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		category := req.FormValue("category")
		strID := req.FormValue("id")
		if category == "" && strID == "" {
			data := models.PhotoDAO.FindFullAll()
			formatter.HTML(w, http.StatusOK, "console/photos/photos", data)
		} else if category == "" {
			id, _ := strconv.Atoi(strID)
			data, has := models.PhotoDAO.FindFullByID(id)
			if !has {
				formatter.Text(w, http.StatusBadRequest, "Photo not found")
				return
			}
			formatter.HTML(w, http.StatusOK, "console/photos/photo", data)
		} else if strID == "" {
			data := models.PhotoDAO.FindFullByCategory(category)
			formatter.HTML(w, http.StatusOK, "console/photos/photos", data)
		}
	}

}
