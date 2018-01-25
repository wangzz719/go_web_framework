package application

import (
	"go-web/src/router"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"go-web/src/middleware"
)

type Application struct {
	router *httprouter.Router
}

func NewApplication(routers []*router.Router) *Application {
	appRouter := httprouter.New()
	for _, r := range routers {
		appRouter.GET(r.Path, middleware.PatchHandle(r.Handler.GET))
		appRouter.POST(r.Path, middleware.PatchHandle(r.Handler.POST))
		appRouter.PUT(r.Path, middleware.PatchHandle(r.Handler.PUT))
		appRouter.PATCH(r.Path, middleware.PatchHandle(r.Handler.PATCH))
		appRouter.HEAD(r.Path, middleware.PatchHandle(r.Handler.HEAD))
		appRouter.OPTIONS(r.Path, middleware.PatchHandle(r.Handler.OPTIONS))
		appRouter.DELETE(r.Path, middleware.PatchHandle(r.Handler.DELETE))
	}
	return &Application{router: appRouter}
}

func (app *Application) ListenAndServe(addr string) {
	err := http.ListenAndServe(":8000", app.router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
