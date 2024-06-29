package services

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"kms/models"
	"kms/repositories"
	"time"
)

var articleRepository *repositories.MongoArticleRepository

type ArticleService struct {
	articleRepositories *repositories.MongoArticleRepository
}

type ArticleServiceInterface interface {
	CreteArticle(title string, content string, authorId string) error
	Update(articleId uint, title string, content string, authorId string) error
	Delete(articleId uint) error
	FindById(articleId uint) (*models.Article, error)
	FindAll() ([]*models.Article, error)
}

func NewArticleService(articleRepo *repositories.MongoArticleRepository) *ArticleService {
	articleRepository = articleRepo
	return &ArticleService{articleRepositories: articleRepo}
}

func (service *ArticleService) CreteArticle(title string, content string, authorId string) error {
	if title == "" {
		return errors.New("title is required")
	}
	if content == "" {
		return errors.New("content is required")
	}
	if authorId == "" {
		return errors.New("authorId is required")
	}

	article := &models.Article{ID: uuid.New().String(), Title: title, Content: content, AuthorId: authorId, CreatedAt: time.Now()}
	err := articleRepository.CreateArticle(article)

	if err != nil {
		fmt.Printf("create article failed, err:%v\n", err)
		return errors.New("create article failed")
	}
	return nil
}

func (service *ArticleService) Update(articleId string, title string, content string) error { // TODO: 检查authorId 是否正确
	if title == "" {
		return errors.New("title is required")
	}
	if content == "" {
		return errors.New("content is required")
	}

	article := &models.Article{ID: articleId, Title: title, Content: content, UpdatedAt: time.Now()}
	err := articleRepository.UpdateArticle(article)

	if err != nil {
		fmt.Printf("update article %s failed, err:%v\n", articleId, err)
		return errors.New("update article failed")
	}
	return nil
}

func (service *ArticleService) Delete(articleId string) error { // TODO: 检查authorId 是否正确
	err := articleRepository.DeleteArticle(articleId)
	if err != nil {
		fmt.Printf("delete article %s failed, err:%v\n", articleId, err)
		return errors.New("delete article failed")
	}
	return nil
}

func (service *ArticleService) FindById(articleId string) (*models.Article, error) {
	var article *models.Article
	article, err := articleRepository.FindArticleById(articleId)
	if err != nil {
		fmt.Printf("find article %s failed, err:%v\n", articleId, err)
		return article, errors.New("find article failed")
	}
	return article, nil
}

func (service *ArticleService) FindAll() ([]*models.Article, error) {
	var articles []*models.Article
	articles, err := articleRepository.FindAllArticle()
	if err != nil {
		fmt.Printf("find all article failed, err:%v\n", err)
		return articles, errors.New("find all article failed")
	}
	return articles, err
}
