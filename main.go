package main

import "go-reload-json/services"

const (
	PORT = ":8080"
)

func main() {
	services.StartServer(PORT)
}
