package api

import (
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

func NewServer() *negroni.Negroni {
	formatter = render.New(render.Options{
		IndentJSON: true,
	})

	jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("SeCrEt"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	n := negroni.Classic()
	router := mux.NewRouter()

	InitRouter(router)

	n.UseHandler(router)

	return n
}

func InitRouter(router *mux.Router) {

}
