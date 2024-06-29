package router

import (
	"github.com/gin-gonic/gin"
	"kms/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/articles", controllers.CreateArticle)
	router.PUT("/articles/:articleId", controllers.UpdateArticle)
	router.DELETE("/articles/:articleId", controllers.DeleteArticle)
	router.GET("/articles", controllers.GetAllArticles)
	router.GET("/articles/:articleId", controllers.GetArticleByID)
	return router
}
