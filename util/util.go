package util

import (
	"crypto/md5"
	"fmt"
	"io"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sunyatsuntobee/server/logger"
)

type Error struct {
	Msg string `json:"error_msg"`
}

type JWTMessage struct {
	Token string `json:"token"`
}

const (
	TimeLayout string = "2006-01-02 15:04"
)

func MD5Hash(value string) string {
	hash := md5.New()
	io.WriteString(hash, value)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func NewJWT(id int, typ int) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":  "ToBEE",
		"iat":  time.Now().Unix(),
		"exp":  time.Now().AddDate(0, 0, 1).Unix(),
		"aud":  id,
		"type": typ,
	})
	signed, err := token.SignedString("secret")
	logger.LogIfError(err)
	return signed
}
