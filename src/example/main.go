package main

import (
	"go-web/src/router"
	"go-web/src/application"
	"go-web/src/example/handlers"
	basehandler "go-web/src/handlers"
)

func main() {
	helloWorldHandler := handlers.HelloWorldHandler{BaseHandler: basehandler.BaseHandler{}}
	routers := []*router.Router{
		{Path: "/hello", Handler: &helloWorldHandler},
	}

	app := application.NewApplication(routers)
	app.ListenAndServe(":8000")
}
