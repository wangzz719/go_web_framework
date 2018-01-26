package middleware

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"runtime/debug"
	"log"
)

func MiddlewareHandle(fn httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(), http.StatusInternalServerError)
				log.Fatalf("WARN: panic in %v - %v", fn, e)
				debug.PrintStack()
			}
		}()
		log.Println("before reqest: ", r.URL.Path) // do something before request
		fn(w, r, params)
		log.Println("after reqest:", r.URL.Path) // do something after request
	}
}
