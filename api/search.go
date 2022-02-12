package api

import (
	"github.com/gin-gonic/gin"
)

type SearchQuery struct {
	Page int `form:"page"`
	Size int `form:"size"`
}

func (api *Api) textSearch(c *gin.Context) {
	var searchQuery SearchQuery

	if err := c.ShouldBindQuery(&searchQuery); err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	list, err := api.esClient.SearchItem(searchQuery.Page, searchQuery.Size)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"data": list,
	})
}

type ItemBody struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

func (api *Api) addItem(c *gin.Context) {
	var body ItemBody
	err := c.ShouldBindJSON(&body)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "No",
		})
		return
	}

	res, err := api.esClient.CreateIndex(body.Id, body.Title)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": res,
	})
}
