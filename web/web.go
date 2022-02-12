package web

import (
	"github.com/gin-gonic/gin"
)

type Web struct {
	app *gin.Engine
}

func New(_app *gin.Engine) *Web {
	return &Web{
		app: _app,
	}
}

func (web *Web) InitRoute() {
	routeWeb := web.app.Group("/")
	{
		routeWeb.GET("/", func(c *gin.Context) {
			c.HTML(200, "index.tmpl", gin.H{
				"title": "Text",
			})
		})
	}
}
