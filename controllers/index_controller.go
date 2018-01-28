package controllers

import (
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func initIndexRouter(router *mux.Router, formatter *render.Render) {
	router.HandleFunc("/", indexHandler(formatter))
	router.HandleFunc("/profile", indexProfileHandler(formatter))
}

func indexHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		html, _ := ioutil.ReadFile("./views/index/photos.html")
		formatter.HTML(w, http.StatusOK, "index", template.HTML(string(html)))
	}

}

func indexProfileHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		html, _ := ioutil.ReadFile("./views/index/profile.html")
		formatter.HTML(w, http.StatusOK, "index", template.HTML(string(html)))
	}

}
