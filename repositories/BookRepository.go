package repositories

import (
	"errors"
	"project-1/database"
	"project-1/models"
)

var (
	err error
)

func AddBook(book models.Book) (models.Book, error) {
	db := database.GetDB()

	err := db.Create(&book).Error

	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func UpdateBook(updatedBook models.Book, id int) (models.Book, error) {
	db := database.GetDB()
	book := models.Book{}

	result := db.Model(&book).Where("id = ?", id).Updates(&updatedBook)

	if result.RowsAffected == 0 {
		return models.Book{}, errors.New("no rows affected")
	}

	if result.Error != nil {
		return models.Book{}, result.Error
	}

	err = db.Model(&book).Where("id = ?", id).First(&updatedBook).Error
	if err != nil {
		return models.Book{}, err
	}

	return updatedBook, nil
}

func DeleteBook(id int) error {
	db := database.GetDB()
	book := models.Book{}

	result := db.Where("id = ?", id).Delete(&book)

	if result.RowsAffected == 0 {
		return errors.New("no rows affected")
	}

	if result.Error != nil {
		return result.Error
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
