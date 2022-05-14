package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"infomap.backend/handlers"
	"infomap.backend/infra"
	"infomap.backend/models"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		panic(err)
	}
	infra.Init()
	db := infra.GetDB()

	if !db.Migrator().HasTable(&models.Article{}) {
		db.AutoMigrate(&models.Article{})
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello,world",
		})
	})
	r.GET("/articles", handlers.Articles)
	r.Run()
}
