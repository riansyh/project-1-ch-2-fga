package routers

import "github.com/gin-gonic/gin"

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/books")
	router.GET("/books/:id")
	router.POST("/books")
	router.PUT("/books")
	router.DELETE("/books")

	return router
}
