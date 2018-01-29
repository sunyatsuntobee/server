package main

import (
	"github.com/sunyatsuntobee/server/controllers"
	"github.com/sunyatsuntobee/server/logger"
)

func main() {
	server := controllers.NewServer()
	server.Run(":8080")
	logger.I.Println("Server is running")
}
