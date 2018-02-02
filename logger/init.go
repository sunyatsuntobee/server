package logger

import (
	"log"
	"os"
)

const IPath string = "./logger/logs/info"
const EPath string = "./logger/logs/error"

var I *log.Logger
var E *log.Logger

func init() {
	truncate()

	infoFile, err := os.OpenFile(IPath, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	I = log.New(infoFile, "[INFO]", log.Ldate|log.Ltime)

	errorFile, err := os.OpenFile(EPath, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	E = log.New(errorFile, "[ERROR]", log.Ldate|log.Ltime)
}

func LogIfError(err error) {
	if err != nil {
		E.Println(err)
	}
}

func truncate() {
	os.Truncate(IPath, 0)
	os.Truncate(EPath, 0)
}
