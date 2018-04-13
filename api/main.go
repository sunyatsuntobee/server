package main

import "github.com/sunyatsuntobee/server/api/api"

func main() {
	server := api.NewServer()
	// port := ":8080"
	port := ":80"
	server.Run(port)
}
