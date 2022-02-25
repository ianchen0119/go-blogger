package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Route struct {
	/* router name */
	Name string
	/* http method */
	Method string
	/* pattern of uri */
	Pattern string
	/* gin handler */
	HandlerFunc gin.HandlerFunc
}

type Routes []Route

func NewRouter() *gin.Engine {
	router := gin.Default()
	group := router.Group("/api")

	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			group.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			group.POST(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			group.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

var routes = Routes{
	{
		"GetPostList",
		http.MethodGet,
		"/list",
		GetPostList,
	},
	{
		"GetContentById",
		http.MethodGet,
		"/post/:id",
		GetContentById,
	},
	{
		"PostContentById",
		http.MethodPost,
		"/post/:id",
		PostContentById,
	},
	{
		"DeleteContentById",
		http.MethodDelete,
		"/post/:id",
		DeleteContentById,
	},
}
