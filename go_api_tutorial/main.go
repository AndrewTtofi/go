package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func addBooks(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	books, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, c.Errors)
		return
	}
	c.IndentedJSON(http.StatusOK, books)
}
func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func checkoutBook(c *gin.Context) {
	book, done := checkBookExists(c)
	if done {
		return
	}
	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
	book, done := checkBookExists(c)
	if done {
		return
	}
	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

func checkBookExists(c *gin.Context) (*book, bool) {
	id, exists := c.GetQuery("id")
	if exists == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter"})
		return nil, true
	}
	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
	}
	return book, false
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", addBooks)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)
	router.Run("localhost:8080")
}
