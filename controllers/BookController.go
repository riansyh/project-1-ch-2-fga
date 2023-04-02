package controllers

import (
	"fmt"
	"net/http"
	"project-1/models"
	"project-1/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateBook(ctx *gin.Context) {
	var newBook models.Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBook, err := repositories.AddBook(newBook)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_status":  "Database Error",
			"error_message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated,
		newBook,
	)
}

func UpdateBook(ctx *gin.Context) {
	bookId := ctx.Param("id")
	idInt, _ := strconv.Atoi(bookId)

	var UpdatedBook models.Book

	if err := ctx.ShouldBindJSON(&UpdatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	UpdatedBook, err := repositories.UpdateBook(UpdatedBook, idInt)

	if err == gorm.ErrRecordNotFound || err.Error() == "no rows affected" {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookId),
		})
		return
	}

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK,
		UpdatedBook,
	)
}

func GetBook(ctx *gin.Context) {
	bookId := ctx.Param("id")
	idInt, _ := strconv.Atoi(bookId)

	var bookData models.Book

	bookData, err := repositories.GetBook(idInt)

	if err == gorm.ErrRecordNotFound {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookId),
		})
		return
	}

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK,
		bookData,
	)
}

func GetAllBooks(ctx *gin.Context) {
	var results = []models.Book{}

	results, err := repositories.GetAllBooks()

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, results)
}

func DeleteBook(ctx *gin.Context) {
	bookId := ctx.Param("id")
	idInt, _ := strconv.Atoi(bookId)

	err := repositories.DeleteBook(idInt)

	if err == gorm.ErrRecordNotFound || err.Error() == "no rows affected" {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookId),
		})
		return
	}

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Book deleted successfully",
	})
}
