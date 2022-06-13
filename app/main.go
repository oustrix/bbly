package main

import (
	"log"

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
	r.POST("/", handlers.Save)
	r.GET("/:id", handlers.Redirect)

	log.Fatal(r.Run(":80"))
}
