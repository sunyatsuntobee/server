package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sunyatsuntobee/server/models"
)

func initConsoleUsersRouter(router *mux.Router) {
	router.HandleFunc("/users", consoleUsersGetHandler()).
		Methods(http.MethodGet)
}

func consoleUsersGetHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if req.FormValue("id") == "" {
			data := models.UserDAO.FindAll()
			formatter.HTML(w, http.StatusOK, "console/users/users", data)
		}
	}

}
