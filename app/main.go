package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"bbly/internal/handlers"
	"bbly/pkg/pg"
)

func main() {
	err := pg.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	r.LoadHTMLGlob("./web/templates/*")

	r.GET("/", handlers.Index)
	//r.PUT("/", handlers.Save)
	r.GET("/:id", handlers.Redirect)

	port := os.Getenv("PORT")
	log.Fatal(r.Run(":" + port))
}
