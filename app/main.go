package main

import (
	"bbly/internal/handlers"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./web/templates/*")

	r.GET("/", handlers.Index)
	r.GET("/:id", handlers.Redirect)

	port := os.Getenv("PORT")
	log.Fatal(r.Run(":" + port))
}
