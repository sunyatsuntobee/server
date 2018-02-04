package controllers

import (
	"html/template"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/sunyatsuntobee/server/api/api"
	"github.com/sunyatsuntobee/server/models"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var (
	formatter     *render.Render
	jwtMiddleware *jwtmiddleware.JWTMiddleware
)

// NewServer initializes and returns a new negroni server
func NewServer() *negroni.Negroni {
	formatter = render.New(render.Options{
		Directory:     "views",
		IndentJSON:    true,
		UnEscapeHTML:  true,
		IsDevelopment: true,
		Funcs: []template.FuncMap{
			template.FuncMap{"UniqueAt": models.UniqueAt},
		},
	})

	jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	n := negroni.Classic()
	router := mux.NewRouter()

	initRouter(router)
	api.InitRouter(router)

	n.UseHandler(router)

	return n
}

func initRouter(router *mux.Router) {

	// Static files
	router.PathPrefix("/static").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("./static/"))))

	// Develop
	router.HandleFunc("/dev", func(w http.ResponseWriter, req *http.Request) {
		formatter.HTML(w, http.StatusOK, "dev", nil)
	})

	initLoginRouter(router)

	initRegisterRouter(router)

	initConsoleRouter(router)
}

func secureHandler(handler http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		jwtMiddleware.HandlerWithNext(w, req, handler)
	}

}
