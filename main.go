package main

import (
	"fmt"
	"net/http"
	"tests/gin-gorm/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err = gorm.Open("postgres", `host=localhost user=bgk dbname=gorm_test sslmode=disable`)
	panicOnError(err)
	db.LogMode(true)
	database.Setup(db)
	defer db.Close()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", index)
	r.GET("/books", booksIndex)
	r.GET("/books/show", bookShow)
	r.GET("/books/create", bookCreateForm)
	r.POST("/books/create/process", bookCreateProcess)
	r.GET("/books/update", bookUpdateForm)
	r.POST("/books/update/process", bookUpdateProcess)
	r.GET("/books/delete/process", bookDeleteProcess)
	r.Run()
}

func index(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/books")
}

func booksIndex(c *gin.Context) {
	bks := database.All(db)
	if bks == nil {
		c.AbortWithStatus(404)
	} else {
		c.HTML(http.StatusOK, "books.gohtml", bks)
	}
}

func bookShow(c *gin.Context) {
	isbn := c.Query("isbn")
	bk := database.GetOne(db, isbn)
	if bk == nil {
		c.AbortWithStatus(404)
	} else {
		c.HTML(http.StatusOK, "show.gohtml", bk)
	}
}

func bookCreateForm(c *gin.Context) {
	c.HTML(http.StatusOK, "create.gohtml", nil)
}

func bookCreateProcess(c *gin.Context) {
	isbn := c.PostForm("isbn")
	title := c.PostForm("title")
	author := c.PostForm("author")
	price := c.PostForm("price")
	if _, err := database.SetOne(db, isbn, title, author, price); err != nil {
		fmt.Println("err")
		c.AbortWithStatus(500)
	}
	c.Redirect(301, "/")
}

func bookUpdateForm(c *gin.Context) {
	isbn := c.Query("isbn")
	bk := database.GetOne(db, isbn)
	if bk == nil {
		c.AbortWithStatus(404)
	} else {
		c.HTML(http.StatusOK, "update.gohtml", bk)
	}
}

func bookUpdateProcess(c *gin.Context) {
	isbn := c.PostForm("isbn")
	title := c.PostForm("title")
	author := c.PostForm("author")
	price := c.PostForm("price")
	if _, err := database.UpdateOne(db, isbn, title, author, price); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(500)
	} else {
		c.Redirect(301, "/")
	}
}

func bookDeleteProcess(c *gin.Context) {
	isbn := c.Query("isbn")
	fmt.Print(isbn)
	if err := database.DeleteOne(db, isbn); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(500)
	} else {
		c.Redirect(301, "/")
	}
}
