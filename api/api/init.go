package api

import (
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/unrolled/render"
)

var (
	formatter     *render.Render
	jwtMiddleware *jwtmiddleware.JWTMiddleware
)

func init() {
	formatter = render.New(render.Options{
		IndentJSON: true,
	})

	jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("SeCrEt"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
}
