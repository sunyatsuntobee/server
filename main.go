package main

import "github.com/sunyatsuntobee/server/controllers"

func main() {
	server := controllers.NewServer()
	server.Run(":8080")
}
