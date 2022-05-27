package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"bbly/pkg/pg"
)

func Redirect(c *gin.Context) {
	id := c.Param("id")
	row := pg.DB.QueryRow(context.Background(), "SELECT url FROM links WHERE id=$1 LIMIT 1;", id)
	var url string
	err := row.Scan(&url)
	if err != nil {
		log.Fatal("problems to select url with id" + id)
	}
	log.Printf("found url %s by id %s\n", url, id)
	c.Redirect(http.StatusMovedPermanently, url)
	c.Abort()
}
