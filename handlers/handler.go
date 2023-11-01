package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/siddharthans2000/go-crud-gin-basic/model"
)

func SetupHandlers(r *gin.Engine) {
	r.GET("/getBooks", getBooks)
	r.GET("/getBook/:id", getBook)
	r.POST("/insertBook", insertBook)
	r.DELETE("/removeBook/:id", removeBook)
	r.PUT("/updateBook", updateBook)
}
func parseToInt(key string) int64 {
	value, err := strconv.ParseInt(key, 0, 0)
	if err != nil {
		log.Fatal("Error with Parsing")
	}
	return value
}
func getBook(c *gin.Context) {
	id := c.Param("id")
	ID := parseToInt(id)
	_, book := model.GetBook(ID)
	c.IndentedJSON(http.StatusAccepted, book)
}

func getBooks(c *gin.Context) {
	books := model.GetBooks()
	c.IndentedJSON(http.StatusCreated, books)
}
func insertBook(c *gin.Context) {
	book := model.Book{}
	if err := c.BindJSON(&book); err != nil {
		log.Fatal("Error with the data")
		return
	}
	da := book.CreateBook()
	c.IndentedJSON(http.StatusCreated, da)
}
func removeBook(c *gin.Context) {
	id := c.Param("id")
	ID := parseToInt(id)
	book := model.DeleteBook(ID)
	c.IndentedJSON(http.StatusAccepted, book)
}
func updateBook(c *gin.Context) {
	id := c.Param("id")
	ID := parseToInt(id)
	db, existingBook := model.GetBook(ID)
	var newBook model.Book
	if err := c.BindJSON(&newBook); err != nil {
		log.Fatalln("Issue with Binding")
	}
	if newBook.Author != "" {
		existingBook.Author = newBook.Author
	}
	if newBook.Name != "" {
		existingBook.Price = newBook.Price
	}
	if newBook.Publication != "" {
		existingBook.Publication = newBook.Publication
	}
	if newBook.Price != 0 {
		existingBook.Price = newBook.Price
	}
	db.Save(&existingBook)
	c.IndentedJSON(http.StatusAccepted, existingBook)
}
