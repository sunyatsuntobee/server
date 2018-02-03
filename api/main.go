package main

import "github.com/sunyatsuntobee/server/api/api"

func main() {
	server := api.NewServer()
	server.Run(":9090")
}
