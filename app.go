package main

import (
	"fmt"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	apiG "github.com/rubenkristian/searchpage/api"
	"github.com/rubenkristian/searchpage/esmodel"
	webG "github.com/rubenkristian/searchpage/web"
)

type Doc struct {
	Title string `json:"title"`
}

func main() {
	// gin.SetMode(gin.ReleaseMode)
	app := gin.Default()

	app.LoadHTMLGlob("templates/*")

	esClient, err := elasticsearch.NewDefaultClient()

	if err != nil {
		panic(fmt.Sprintf("Elasticsearch Error", err))
	}

	es := esmodel.New(esClient)

	apiGroup := apiG.New(app, es)
	apiGroup.InitRoute()

	webGroup := webG.New(app)
	webGroup.InitRoute()

	app.Run()
}
