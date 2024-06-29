package repositories

import "kms/models"

type ArticleRepository interface {
	CreateArticle(article *models.Article) error
	DeleteArticle(article *models.Article) error
	UpdateArticle(article *models.Article) error
	FindArticleById(id string) (*models.Article, error)
	FindAllArticle() ([]*models.Article, error)
}
