package main

import (
	"go-web/src/router"
	"go-web/src/application"
	"go-web/src/example/handlers"
	basehandler "go-web/src/handlers"
)

func main() {
	routers := []*router.Router{
		{Path: "/hello", Handler: &handlers.HelloWorldHandler{BaseHandler: basehandler.BaseHandler{}}},
	}

	app := application.NewApplication(routers)
	app.ListenAndServe(":8000")
}
