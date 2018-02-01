package controllers

import (
	"net/http"

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
		if req.FormValue("category") == "" {
			data := models.PhotoDAO.FindAll()
			formatter.HTML(w, http.StatusOK, "console/photos/photos", data)
		}
	}

}
