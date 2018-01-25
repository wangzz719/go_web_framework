package handlers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"io"
)

type HanderInterface interface {
	GET(http.ResponseWriter, *http.Request, httprouter.Params)
	POST(http.ResponseWriter, *http.Request, httprouter.Params)
	PUT(http.ResponseWriter, *http.Request, httprouter.Params)
	DELETE(http.ResponseWriter, *http.Request, httprouter.Params)
	OPTIONS(http.ResponseWriter, *http.Request, httprouter.Params)
	HEAD(http.ResponseWriter, *http.Request, httprouter.Params)
	PATCH(http.ResponseWriter, *http.Request, httprouter.Params)
}

type BaseHandler struct {
}

func (h *BaseHandler) GET(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	io.WriteString(w, "method not allowed")
}
func (h *BaseHandler) POST(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	io.WriteString(w, "method not allowed")
}
func (h *BaseHandler) PUT(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	io.WriteString(w, "method not allowed")
}
func (h *BaseHandler) DELETE(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	io.WriteString(w, "method not allowed")
}
func (h *BaseHandler) OPTIONS(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	io.WriteString(w, "method not allowed")
}

func (h *BaseHandler) HEAD(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	io.WriteString(w, "method not allowed")
}

func (h *BaseHandler) PATCH(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	io.WriteString(w, "method not allowed")
}
