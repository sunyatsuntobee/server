package controllers

import (
	"html/template"
	"io/ioutil"
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
		html, _ := ioutil.ReadFile("./views/register.html")
		page := layout{
			Title:   "注册 - 图蜂后台管理系统",
			Content: template.HTML(string(html)),
		}
		formatter.HTML(w, http.StatusOK, "layout", page)
	}

}

func registerPostHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		user := models.User{
			Phone:    req.FormValue("phone"),
			Password: req.FormValue("password"),
			Location: req.FormValue("location"),
		}
		models.UserDAO.InsertOne(&user)
		formatter.JSON(w, http.StatusCreated, nil)
	}

}
