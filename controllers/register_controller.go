package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sunyatsuntobee/server/models"
)

func initRegisterRouter(router *mux.Router) {
	router.HandleFunc("/register", registerGetHandler()).Methods(http.MethodGet)
	router.HandleFunc("/register", registerPostHandler()).
		Methods(http.MethodPost)
}

func registerGetHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.HTML(w, http.StatusOK, "register", nil)
	}

}

func registerPostHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		user := models.User{
			Phone:    req.FormValue("phone"),
			Password: req.FormValue("password"),
		}
		models.UserDAO.InsertOne(&user)
		formatter.JSON(w, http.StatusCreated, nil)
	}

}
