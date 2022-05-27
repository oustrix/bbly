package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"bbly/pkg/pg"
)

func Redirect(c *gin.Context) {
	id := getID(c)
	url, err := searchURL(id)
	if err != nil {
		log.Println("problems to select url with id" + id)
		c.HTML(http.StatusInternalServerError, "server_error.html", gin.H{})
	} else {
		c.Redirect(http.StatusMovedPermanently, url)
		c.Abort()
	}
}

func getID(c *gin.Context) string {
	return c.Param("id")
}

func searchURL(id string) (string, error) {
	row := pg.DB.QueryRow(context.Background(), "SELECT url FROM links WHERE id=$1 LIMIT 1;", id)
	var url string
	err := row.Scan(&url)
	return url, err
}
