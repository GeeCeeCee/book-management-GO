package models

import (
	"fmt"

	"github.com/GeeCeeCee/bookings/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct { // this will create a table name in plural, eg. "books"
	// gorm.Model // this will add extra columns to the model, like CreatedDate, UpdatedDate and DeletedDate
	ID int64 `json:"id"` 
	Name string `json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init () {
	config.Connect()
	db = config.GetDB()

	db.AutoMigrate(&Book{}) //AutoMigrate run auto migration for given models, will only add missing fields, won't delete/change current data
}

func (b *Book) CreateBook() *Book {
	// db.NewRecord(b)
	db.Create(b)
	fmt.Println("Book created", b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookByID(ID int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID = ?", ID).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) *Book {
	var book Book
	db.Where("ID = ?", ID).Delete(&book)
	return &book
}