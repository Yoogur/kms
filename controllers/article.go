package controllers

import (
	"github.com/gin-gonic/gin"
	"kms/models"
	"kms/services"
)

var articleService *services.ArticleService

func NewArticleController(service *services.ArticleService) {
	articleService = service
}

func CreateArticle(c *gin.Context) {
	title := c.Query("title")
	content := c.Query("content")
	authorId := c.Query("authorId")

	err := articleService.CreteArticle(title, content, authorId)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
	
	c.JSON(200, gin.H{"msg": "success"})
}

func UpdateArticle(c *gin.Context) {
	articleId := c.Param("articleId")
	title := c.Query("title")
	content := c.Query("content")

	err := articleService.Update(articleId, title, content)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, gin.H{"msg": "success"})

}

func DeleteArticle(c *gin.Context) {
	articleId := c.Param("articleId")

	err := articleService.Delete(articleId)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(200, gin.H{"msg": "success"})
}

func GetArticleByID(c *gin.Context) {
	articleId := c.Param("articleId")
	var article *models.Article
	article, err := articleService.FindById(articleId)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": article})
}

func GetAllArticles(c *gin.Context) {
	var articles []*models.Article
	articles, err := articleService.FindAll()
	if err != nil {
		c.JSON(400, gin.H{})
		return
	}

	c.JSON(200, gin.H{"data": articles})
}
