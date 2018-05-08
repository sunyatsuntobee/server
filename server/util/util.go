package util

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"image/png"
	"io"
	"mime/multipart"
	"os"
	"strings"
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
	signed, err := token.SignedString([]byte("abcd"))
	logger.LogIfError(err)
	return signed
}

func SaveMultipartFile(path string, file multipart.File) {
	var buf bytes.Buffer
	io.Copy(&buf, file)
	target, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer target.Close()
	target.Write(buf.Bytes())
	logger.LogIfError(err)
}

func SaveBase64AsPNG(code string, path string) {
	code = code[strings.IndexByte(code, ',')+1:]
	unbased, err := base64.StdEncoding.DecodeString(code)
	logger.LogIfError(err)
	r := bytes.NewReader(unbased)
	img, err := png.Decode(r)
	logger.LogIfError(err)
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, os.ModePerm)
	logger.LogIfError(err)
	png.Encode(file, img)
}
