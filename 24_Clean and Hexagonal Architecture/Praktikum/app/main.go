package main

import (
	"Praktikum/infrastructures/http/server"
)

func main() {
	app := server.Server()
	app.Start(":8100")
}
