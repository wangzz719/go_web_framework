package router

import (
	"go-web/src/handlers"
)

type Router struct {
	Path    string
	Handler handlers.HanderInterface
}
