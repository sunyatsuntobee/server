package controllers

import (
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func initRegisterRouter(router *mux.Router, formatter *render.Render) {
	router.HandleFunc("/register", registerHandler(formatter))
}

func registerHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		html, _ := ioutil.ReadFile("./views/register.html")
		page := layout{
			Title:   "注册 - 图蜂后台管理系统",
			Content: template.HTML(string(html)),
		}
		formatter.HTML(w, http.StatusOK, "layout", page)
	}

}
