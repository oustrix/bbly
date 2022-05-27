package handlers

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/dchest/uniuri"
	"github.com/gin-gonic/gin"

	"bbly/pkg/pg"
)

func Save(c *gin.Context) {
	url := getURL(c)

	// if for some full url shorten url already exist it's just response with already existing shorten url
	isExist, shortURL := checkForExisting(url)
	if isExist {
		responseDone(c, shortURL)
		return
	}

	// generate and saving shorten url
	shortURL = randomUrl()
	log.Printf("Generate shorten URL: %s for %s", shortURL, url)
	_, err := pg.DB.Exec(context.Background(), "INSERT INTO links (id, url, visits) VALUES ($1, $2, $3)", shortURL, url, 0)
	if err != nil {
		responseServerError(c, err)
	} else {
		log.Printf("%s saved in DB", shortURL)
		responseDone(c, shortURL)
	}

}

// get URL from PostForm
func getURL(c *gin.Context) string {
	return c.PostForm("url")
}

func checkForExisting(url string) (bool, string) {
	var id string
	row := pg.DB.QueryRow(context.Background(), "SELECT id FROM links WHERE url=$1", url)
	err := row.Scan(&id)
	if err != nil {
		return false, ""
	}
	return true, id
}

// generates random short URL
func randomUrl() string {
	var newURL string
	isExist := true
	rand.Seed(time.Now().UnixNano())
	for isExist {
		newURL = uniuri.NewLen(6)
		row := pg.DB.QueryRow(context.Background(), "SELECT id FROM links WHERE url=$1 LIMIT 1", newURL)
		err := row.Scan()
		if err != nil {
			isExist = false
		}
	}
	return newURL
}

func responseServerError(c *gin.Context, err error) {
	c.HTML(http.StatusInternalServerError, "server_error.html", gin.H{})
	log.Println(err)
}

func responseDone(c *gin.Context, shortURL string) {
	c.HTML(http.StatusOK, "done.html", gin.H{
		"shortURL": shortURL,
	})
}
