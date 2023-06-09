package route

import (
	"log"
	"net/http"
)

type HandlerFunc func(ctx *Context)

type Ruu struct {
	router *router
	middle []HandlerFunc
}

func (r *Ruu) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	c := newContext(writer, request)
	c.handlers = append(c.handlers, r.middle...)
	r.router.handle(c)
}

func (r *Ruu) Use(middlewares ...HandlerFunc) {
	r.middle = append(r.middle, middlewares...)
}

func (r *Ruu) GET(path string, handler HandlerFunc) {
	r.addRoute("GET", path, handler)
}

func (r *Ruu) POST(path string, handler HandlerFunc) {
	r.addRoute("POST", path, handler)
}

func (r *Ruu) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	r.router.addRoute(method, pattern, handler)
}

func (r *Ruu) Run(addr string) (err error) {
	return http.ListenAndServe(addr, r)
}

func New() *Ruu {
	return &Ruu{router: newRouter()}
}
