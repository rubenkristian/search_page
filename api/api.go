package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rubenkristian/searchpage/esmodel"
)

type Api struct {
	app      *gin.Engine
	esClient *esmodel.EsModel
}

func New(_app *gin.Engine, esClient *esmodel.EsModel) *Api {
	return &Api{
		app:      _app,
		esClient: esClient,
	}
}

func (api *Api) InitRoute() {
	routeApi := api.app.Group("/api")
	{
		routeApi.GET("/search", api.textSearch)
		routeApi.POST("/item", api.addItem)
	}
}
