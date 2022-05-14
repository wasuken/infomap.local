package services

import (
	"infomap.backend/infra"
	"infomap.backend/models"
)

func GetArticles() []models.Article {
	db := infra.GetDB()
	var articles []models.Article
	db.Find(&articles)
	return articles
}

func SearchArticles(q string, lat, lon float64) []models.Article {
	db := infra.GetDB()
	var articles []models.Article
	db.Find(&articles)
	return articles
}
