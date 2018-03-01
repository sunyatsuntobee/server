package main

import (
	"github.com/sunyatsuntobee/server/controllers"
	"github.com/sunyatsuntobee/server/logger"
)

func main() {
	server := controllers.NewServer()
	port := ":8080"
	// port := ":80"
	server.Run(port)
	logger.I.Println("Web Server is running")
}
