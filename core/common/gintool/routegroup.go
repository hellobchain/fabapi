package gintool

import "github.com/gin-gonic/gin"

type RouteGroup struct {
	Prefix  string  //路由分组前缀
	Route   []Route //子路由
	Comment string  //模块注释信息
	Module  string  //模块名称
}

type Route struct {
	Method      string            //http请求方法
	Path        string            //请求路径
	HandlerFunc []gin.HandlerFunc //调用方法
	Comment     string            //路由注释信息
}

func NewRoute(method, path string, handler ...gin.HandlerFunc) Route {
	return Route{
		Method:      method,
		Path:        path,
		HandlerFunc: handler,
	}
}

//给当前路由添加注释信息
func (r Route) AddComment(comment string) Route {
	r.Comment = comment
	return r
}

//将分组后的路由集合根据分组注册
func AddRouteGroups(engine *gin.Engine, prefix string, routes []RouteGroup) {
	api := engine.Group(prefix)
	for _, route := range routes {
		group := api.Group(route.Prefix)
		for _, r := range route.Route {
			group.Handle(r.Method, r.Path, r.HandlerFunc...)
		}
	}
}
