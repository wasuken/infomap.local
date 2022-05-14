package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"infomap.backend/services"
)

func Articles(c *gin.Context) {
	articles := services.GetArticles()
	c.JSON(http.StatusOK, gin.H{
		"articles": articles,
	})
}

func SearchArticles(c *gin.Context) {
	// 検索クエリ
	q := c.GetString("q")
	// 座標
	lat := c.GetFloat64("lat")
	lon := c.GetFloat64("lon")
	articles := services.SearchArticles(q, lat, lon)
	c.JSON(http.StatusOK, gin.H{
		"articles": articles,
	})
}
