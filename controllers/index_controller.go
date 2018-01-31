package controllers

import (
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func initIndexRouter(router *mux.Router) {
	// router.HandleFunc("/", secureHandler(indexHandler()))
	router.HandleFunc("/", indexHandler())
	router.HandleFunc("/profile", indexProfileHandler())
}

func indexHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		html, _ := ioutil.ReadFile("./views/index/photos.html")
		formatter.HTML(w, http.StatusOK, "index/photos", template.HTML(string(html)))
	}

}

func indexProfileHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		html, _ := ioutil.ReadFile("./views/index/profile.html")
		formatter.HTML(w, http.StatusOK, "index/profile", template.HTML(string(html)))
	}

}
