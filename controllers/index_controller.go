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
	router.HandleFunc("/photolives", indexPhotoLivesHandler())
}

func indexHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		html, _ := ioutil.ReadFile("./views/index/photos.html")
		formatter.HTML(w, http.StatusOK, "index", template.HTML(string(html)))
	}

}

func indexProfileHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		html, _ := ioutil.ReadFile("./views/index/profile.html")
		formatter.HTML(w, http.StatusOK, "index", template.HTML(string(html)))
	}

}

func indexPhotoLivesHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		html, _ := ioutil.ReadFile("./views/index/photo_lives.html")
		formatter.HTML(w, http.StatusOK, "index", template.HTML(string(html)))
	}

}
