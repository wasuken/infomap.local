package services

import (
	"log"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"infomap.backend/infra"
	"infomap.backend/models"
)

func setup() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("failed load .env")
	}
	infra.InitTest()
	db := infra.GetDB()
	if !db.Migrator().HasTable(&models.Article{}) {
		db.AutoMigrate(&models.Article{})
	}

	gin.SetMode(gin.TestMode)
}

func teardown() {
	if _, err := os.Stat("../test.db"); err == nil {
		r := os.Remove("../test.db")
		if r != nil {
			panic(r)
		}
	}
}

func TestMain(m *testing.M) {
	setup()
	ret := m.Run()
	if ret == 0 {
		teardown()
	}
	os.Exit(ret)
}

func TestGetArticles(t *testing.T) {
	articles := GetArticles()
	expect := []models.Article{}
	assert.Equal(t, expect, articles)
}
