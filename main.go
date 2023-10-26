package main

import (
	"mini-project-golang/config"
	"mini-project-golang/route"
)

func main() {
	config.InitDB()

	e := route.NewRoute()

	e.Logger.Fatal(e.Start(":8181"))
}