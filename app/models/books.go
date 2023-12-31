package models

type Book struct {
	Title  string `json:"title" gorm:"column:title;type:varchar(255);not null"`
	Author string `json:"author" gorm:"column:author;type:varchar(255);not null"`
	BaseModel
}

func (Book) TableName() string {
	return "tbl_books"
}
