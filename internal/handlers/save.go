package handlers

import (
	"context"
	"math/rand"
	"time"

	"github.com/dchest/uniuri"
	"github.com/gin-gonic/gin"

	"bbly/pkg/pg"
)

func Save(c *gin.Context) error {
	url := c.PostForm("url")
	rand.Seed(time.Now().UnixNano())
	shortURL := randomUrl()
	_, err := pg.DB.Exec(context.Background(), "INSERT INTO links (id, url, visits) VALUES ($1, $2, $3)", shortURL, url, 0)
	if err != nil {
		return err
	}

	return nil
}

// generates random short URL
func randomUrl() string {
	var newURL string
	isExist := true
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
