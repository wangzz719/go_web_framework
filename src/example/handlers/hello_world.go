package handlers

import (
	"go-web/src/handlers"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"io"
)

type HelloWorldHandler struct {
	handlers.BaseHandler
}

func (h *HelloWorldHandler) GET(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	io.WriteString(w, "hello world")
}
