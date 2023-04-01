package repositories

import (
	"project-1/database"
	"project-1/models"
)

var (
	err error
)

func AddBook(book models.Book) error {
	db := database.GetDB()

	err := db.Create(&book).Error

	if err != nil {
		return err
	}

	return nil
}

func UpdateBook(updatedBook models.Book, id int) error {
	db := database.GetDB()
	book := models.Book{}

	err := db.Model(&book).Where("id = ?", id).Updates(updatedBook).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteBook(id int) error {
	db := database.GetDB()
	book := models.Book{}

	err := db.Where("id = ?", id).Delete(&book).Error
	if err != nil {
		return err
	}

	return nil
}

func GetBook(id int) (models.Book, error) {
	db := database.GetDB()

	book := models.Book{}

	err := db.First(&book, "id = ?", id).Error

	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func GetAllBooks() ([]models.Book, error) {
	db := database.GetDB()
	books := []models.Book{}

	db.Find(&books)

	if err != nil {
		return []models.Book{}, err
	}

	return books, nil
}
