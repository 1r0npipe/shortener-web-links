package handler

import (
	_ "fmt"
	"net/http"

	"github.com/1r0npipe/shortener-web-links/internal/generator"
	"github.com/1r0npipe/shortener-web-links/internal/model"
	"github.com/1r0npipe/shortener-web-links/internal/storage"
	"github.com/gin-gonic/gin"
)

type Link struct {
	URL    *string `json:"url"`
	UserID *string `json:"userID"`
	TTL    *int    `json:"ttl,omitempty"`
}


func CheckHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}



func GenerateNewLink(c *gin.Context) {
	var link *Link
	var m model.Item
	redisCl := storage.ClientRedis{}
	redisDB, err := redisCl.New(storage.DefaultOptions)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, "Can't get body from request")
		return
	}

	err = c.BindJSON(&link)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Can't parse JSON from body")
		return
	}
	shortURL, err := generator.GenerateShortUrl(*link.URL, *link.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	m.FullLink = *link.URL
	m.ShortLink = shortURL
	m.TTL = uint(*link.TTL)
	m.UserID = *link.UserID
	m.Count = 0
	err = redisDB.Put(shortURL, m)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	// c.JSON(http.StatusOK, gin.H{
	// 	"got_url": link.URL,
	// 	"ttl":     link.TTL,
	// })
	c.JSON(http.StatusOK, "Link has been generated")
	return
}

func RedirectByShortUrl(c *gin.Context) {
	// shortUrl := c.Param("shortUrl")
	// itirialUrl := retriveUrl(shortUrl)
	// c.Redirect(302, initialUrl)
}

func GetStatById(c *gin.Context) {

}
