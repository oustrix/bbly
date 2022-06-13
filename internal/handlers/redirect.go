package handlers

// TODO: add tests for searchURLAndVisits
// TODO: add tests for Redirect

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"bbly/pkg/pg"
)

type Row struct {
	id     string
	url    string
	visits int
}

func Redirect(c *gin.Context) {
	var row Row
	var err error

	row.id = getID(c)
	row.url, row.visits, err = searchURLAndVisits(row.id)
	if err != nil {
		log.Println("problems to select url with id" + row.id)
		c.HTML(http.StatusInternalServerError, "server_error.html", gin.H{})
	} else {
		_, err := pg.DB.Exec(context.Background(), "UPDATE links SET visits=$1 WHERE id=$2", row.visits+1, row.id)
		if err != nil {
			log.Printf("failed to increment visits counter for %s\n", row.id)
		}
		c.Redirect(http.StatusMovedPermanently, row.url)
		c.Abort()
	}
}

// get id from request
func getID(c *gin.Context) string {
	return c.Param("id")
}

// search for full URL and visits from DB by id
func searchURLAndVisits(id string) (string, int, error) {
	row := pg.DB.QueryRow(context.Background(), "SELECT url, visits FROM links WHERE id=$1 LIMIT 1;", id)
	var url string
	var visits int
	err := row.Scan(&url, &visits)
	return url, visits, err
}
