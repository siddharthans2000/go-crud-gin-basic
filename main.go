package main

import (
	"github.com/gin-gonic/gin"
	"github.com/siddharthans2000/go-crud-gin-basic/handlers"
)

func main() {
	r := gin.Default()
	handlers.SetupHandlers(r)
	r.Run("localhost:8080")
}
