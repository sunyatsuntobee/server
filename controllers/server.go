package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

// NewServer initializes and returns a new negroni server
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		Directory:    "views/templates",
		IndentJSON:   true,
		UnEscapeHTML: true,
	})

	n := negroni.Classic()
	router := mux.NewRouter()

	initRouter(router, formatter)

	n.UseHandler(router)

	return n
}

func initRouter(router *mux.Router, formatter *render.Render) {

	// Static files
	router.PathPrefix("/static").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("./static/"))))

	initLoginRouter(router, formatter)

	initRegisterRouter(router, formatter)

	initIndexRouter(router, formatter)
}
