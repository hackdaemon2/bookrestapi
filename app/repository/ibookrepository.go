package repository

import "restapi/app/models"

type IBookRepository interface {
	FindAll(page int, size int) ([]models.Book, error)
	FindOne(id int64) (models.Book, error)
	Create(book models.Book) (models.Book, error)
	CountAll() (int64, error)
	Update(book models.Book) (models.Book, error)
	Delete(book *models.Book) (int64, error)
}
