package routers

import (
	"project-1/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/books", controllers.GetAllBooks)
	router.GET("/books/:id", controllers.GetBook)
	router.POST("/books", controllers.CreateBook)
	router.PUT("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	return router
}
