package handlers

import (
	"github.com/gin-gonic/gin"
)

func SetupHandlers(r *gin.Engine) {
	r.GET("/getAllBooks", getBooks)
	r.GET("/getBook", getBook)
	r.POST("/insertBook", insertBook)
	r.DELETE("/removeBook", removeBook)
	r.PUT("/updateBook", updateBook)
}
func getBook(c *gin.Context) {
	id := c.Param("id")
}

func getBooks(c *gin.Context) {

}
func insertBook(c *gin.Context) {

}
func removeBook(c *gin.Context) {
	id := c.Param("id")
}
func updateBook(c *gin.Context) {
	id := c.Param("id")
}
