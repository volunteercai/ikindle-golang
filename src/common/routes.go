package common

import "github.com/gin-gonic/gin"

var Routes []*Route

type Route struct {
	Path     string
	Children []struct {
		HttpMethod string
		Path       string
		Handler    gin.HandlerFunc
	}
}

func AddRootRoute(rootRoute *Route) {
	Routes = append(Routes, rootRoute)
}

func (c *Route) AddChildRoute(httpMethod string, path string, handler gin.HandlerFunc) {
	c.Children = append(c.Children, struct {
		HttpMethod string
		Path       string
		Handler    gin.HandlerFunc
	}{HttpMethod: httpMethod, Path: path, Handler: handler})
}

type Api struct {
	RouteLoader
	R *Route
}

type RouteLoader interface {
	LoadRoute()
}
