package controllers

import (
	"net/http"
	"strconv"

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
		} else {
			id, _ := strconv.Atoi(req.FormValue("id"))
			data, has := models.UserDAO.FindByID(id)
			if !has {
				formatter.Text(w, http.StatusBadRequest, "User not found")
				return
			}
			formatter.HTML(w, http.StatusOK, "console/users/user", data)
		}
	}

}
