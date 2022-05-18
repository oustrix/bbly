package main

import (
	"bbly/internal/handlers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", handlers.Index)
	r.GET("/:id", handlers.Redirect)

	port := os.Getenv("PORT")
	log.Fatal(r.Run(":" + port))
}
