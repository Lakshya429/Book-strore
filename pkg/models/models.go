package models


import (
	"github/Lakshya429/postgress/pkg/db"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)
type Book struct {
	gorm.Model 
	ID int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Rating int `json:"rating"`
}


func Setup() {
	db.Connect()
	DB = db.GetDB()
	DB.AutoMigrate(&Book{})
}

func CreateBook(book *Book) *Book {
	DB.Create(book)
	return book
}

func GetAllBooks() []Book {
	var books []Book 
	DB.Find((&books))
    return books
}
func GetBook(ID int) *Book {
	book := &Book{}
	DB.Find(book , ID)
	return book
}
func DeleteBook(ID int) *Book {
     DB.Delete(&Book{}, ID) 
	return nil 
}

func UpdateBook( ID int , book *Book) *Book {
	DB.Model(&Book{}).Where("id = ?", ID).Updates(book)
	return book
}