package database

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
)

// Person JSON object
type book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

// Connect to Database and AutoMigrate
func Setup(db *gorm.DB) {

	db.AutoMigrate(&book{})

	b1 := book{Isbn: "978-12391239", Title: "The Sun Also Rises", Author: "Ernest Hemmingway", Price: 9.99}
	b2 := book{Isbn: "978-1503261969", Title: "Emma", Author: "Jayne Austen", Price: 9.44}

	db.Delete(&b1)
	db.Delete(&b2)
	db.Create(&b1)
	db.Create(&b2)
}

// GetAll retrieves all people from database
func All(db *gorm.DB) interface{} {
	var books []book
	if err := db.Find(&books).Error; err != nil {
		fmt.Println(err)
	}
	return books
}

// GetOne retrieves one book from database
func GetOne(db *gorm.DB, isbn string) interface{} {
	bk := book{}
	if err := db.Where("isbn = ?", isbn).First(&bk).Error; err != nil {
		return err
	}
	return bk
}

// SetOne create one new book
func SetOne(db *gorm.DB, isbn string, title string, author string, price string) (interface{}, error) {
	bk, err := checkInput(isbn, title, author, price)
	if err != nil {
		return nil, err
	}
	if err := db.Create(&bk); err != nil {
		return nil, err.Error
	}
	return bk, nil
}

// UpdateOne updates one row
func UpdateOne(db *gorm.DB, isbn string, title string, author string, price string) (interface{}, error) {
	bk, err := checkInput(isbn, title, author, price)
	if err != nil {
		return nil, err
	}
	if err := db.Table("books").Where("isbn = ?", isbn).Updates(&bk); err != nil {
		return nil, err.Error
	}
	return &bk, nil
}

// DeleteOne deletes on book from database
func DeleteOne(db *gorm.DB, isbn string) error {
	fmt.Println(isbn)
	err := db.Table("books").Where("isbn = ?", isbn).Delete(&book{})
	fmt.Print(err)
	return err.Error
}

func checkInput(isbn string, title string, author string, price string) (*book, error) {
	if isbn == "" || title == "" || author == "" || price == "" {
		return nil, errors.New("Error: Please fill out all information")
	}
	f64, err := strconv.ParseFloat(price, 32)
	if err != nil {
		return nil, errors.New("Error: Price must be a number")
	}
	return &book{Isbn: isbn, Title: title, Author: author, Price: float32(f64)}, nil
}
