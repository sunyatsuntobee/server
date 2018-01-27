package controllers

import (
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func initLoginRouter(router *mux.Router, formatter *render.Render) {
	router.HandleFunc("/login", loginHandler(formatter))
}

func loginHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		html, err := ioutil.ReadFile("./views/login.html")
		if err != nil {
			panic(err)
		}
		formatter.HTML(w, http.StatusOK, "layout", template.HTML(string(html)))
	}

}
