package api

import (
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/unrolled/render"
	"strings"
	"net/http"
)

// JSON is a standard form of response data
type JSON struct {
	Status  string      `json:"status"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

var (
	formatter     *render.Render
	jwtMiddleware *jwtmiddleware.JWTMiddleware
)

const (
	dateLayout string = "06-01-02"
	resDir     string = "/root/tobee"
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

// NewJSON creates a new JSON object
func NewJSON(status, msg string, data interface{}) *JSON {
	return &JSON{
		Status:  status,
		Message: msg,
		Data:    data,
	}
}

//用于从http报文中提取token的字符串形式
func getTokenString(req *http.Request) string {
	return strings.Split(req.Header.Get("Authorization"), " ")[1]
}

func handlerSecure(handler http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		err := jwtMiddleware.CheckJWT(w, req)
		if err != nil {
			formatter.JSON(w, http.StatusUnauthorized, NewJSON(
				"认证中间件失败", "Permission denied", nil))
			return
		}
		if handler != nil {
			handler(w, req)
		}
	}

}