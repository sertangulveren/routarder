package routarder

import (
	"fmt"
	"net/http"
	"os"
	"path"
)

type Router struct {
	prefix   string
	routes   []string
	children []*Router
}

func New() *Router {
	return &Router{prefix: "/"}
}

func Serve(addr string) {
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (router *Router) Group(pattern string) *Router {
	r := New()
	r.prefix = path.Join(router.prefix, pattern)
	router.children = append(router.children, r)
	return r
}

func (router *Router) RouteTree() string {
	res := router.selfRoutes()
	for _, c := range router.children {
		res += c.RouteTree()
	}
	return res
}

func (router *Router) selfRoutes() string {
	res := ""
	for _, route := range router.routes {
		res = fmt.Sprintf("%s%s\n", res, path.Join(router.prefix, route))
	}
	return res
}

func (router *Router) HandleFunc(pattern string, handler http.HandlerFunc) {
	router.routes = append(router.routes, pattern)
	http.HandleFunc(path.Join(router.prefix, pattern), handler)
}
