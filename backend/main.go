package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"infomap.backend/models"
)

func main() {
	err := godotenv.Load("./.env")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)
	b.AutoMigrate(&models.Article{})

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello,world",
		})
	})
	r.GET("/articles", func(c *gin.Context) {
		var articles []models.Article
		db.Find(&articles)
		c.JSON(http.StatusOK, gin.H{
			"articles": articles,
		})
	})
}
