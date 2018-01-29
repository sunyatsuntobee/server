package controllers

import (
	"net/http"

	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
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
		Directory:    "views",
		IndentJSON:   true,
		UnEscapeHTML: true,
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

	n.UseHandler(router)

	return n
}

func initRouter(router *mux.Router) {

	// Static files
	router.PathPrefix("/static").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("./static/"))))

	initLoginRouter(router)

	initRegisterRouter(router)

	initIndexRouter(router)
}

func secureHandler(handler http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		jwtMiddleware.HandlerWithNext(w, req, handler)
	}

}
