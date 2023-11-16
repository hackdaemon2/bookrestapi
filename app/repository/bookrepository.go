package repository

import (
	"gorm.io/gorm"
	"restapi/app/models"
	"time"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db}
}

// FindAll fetches all the books in the DB
func (bookRepository *BookRepository) FindAll(page int, size int) ([]models.Book, error) {
	var books []models.Book
	offSet := (page - 1) * size
	bookRepository.db.
		Where("deleted = ?", models.NotDeleted).
		Order("created_at desc").
		Limit(size).
		Offset(offSet).
		Find(&books)
	return books, nil
}

// FindOne retrieves a book by ID from the database
func (bookRepository *BookRepository) FindOne(id int64) (models.Book, error) {
	var book models.Book
	bookRepository.db.
		Where("deleted = ? AND id = ?", models.NotDeleted, id).
		Find(&book)
	return book, nil
}

// Create inserts a new book into the database
func (bookRepository *BookRepository) Create(book models.Book) (models.Book, error) {
	bookRepository.db.Create(&book)
	return book, nil
}

// CountAll all rows in the DB
func (bookRepository *BookRepository) CountAll() (int64, error) {
	var count int64
	bookRepository.db.
		Model(models.Book{}).
		Where("deleted = ?", models.NotDeleted).
		Count(&count)
	return count, nil
}

// Update updates a user in the database
func (bookRepository *BookRepository) Update(book models.Book) (models.Book, error) {
	bookRepository.db.Updates(&book)
	return book, nil
}

// Delete removes a user from the database by ID
func (bookRepository *BookRepository) Delete(book *models.Book) (int64, error) {
	var now = time.Now()
	book.Deleted = 1
	book.DeletedAt = &now
	result := bookRepository.db.Updates(&book)
	return result.RowsAffected, nil
}
