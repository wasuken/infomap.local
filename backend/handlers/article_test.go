package handlers

import (
	"log"
	"net/http/httptest"
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

func TestArticles(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	r.GET("/articles", Articles)

	req := httptest.NewRequest("GET", "/articles", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestSearchArticles(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	r.GET("/search/articles", SearchArticles)

	req := httptest.NewRequest("GET", "/search/articles", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
