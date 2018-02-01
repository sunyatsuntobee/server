package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func initConsolePhotosRouter(router *mux.Router) {
	router.HandleFunc("/photos", consolePhotosGetHandler()).
		Methods(http.MethodGet)
}

func consolePhotosGetHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if req.FormValue("category") == "" {
			formatter.HTML(w, http.StatusOK, "console/photos/photos", nil)
		}
	}

}
