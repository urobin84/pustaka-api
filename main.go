package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"
)

func main() {

	dsn := "root:root@tcp(127.0.0.1:3306)/pustaka_api?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", bookHandler.RootHandler)
	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/book/:id", bookHandler.GetBook)
	v1.POST("/books", bookHandler.CreateBook)
	v1.PUT("/book/:id", bookHandler.UpdateBook)
	v1.DELETE("/book/:id", bookHandler.DeleteBook)

	err = router.Run()
	if err != nil {
		return
	}
}
