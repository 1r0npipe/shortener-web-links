package handler

import (
	"net/http"
	"sync/atomic"

	"github.com/1r0npipe/shortener-web-links/internal/generator"
	"github.com/1r0npipe/shortener-web-links/internal/model"
	"github.com/1r0npipe/shortener-web-links/internal/storage"
	"github.com/gin-gonic/gin"
)

type Link struct {
	URL    *string `json:"url"`
	UserID *string `json:"userID"`
	TTL    *uint   `json:"ttl,omitempty"`
}

func CheckHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"msg":    "Service is alive",
	})
}

func GenerateNewLink(c *gin.Context) {
	var link *Link
	var m model.Item
	redisCl := storage.ClientRedis{}
	redisDB, err := redisCl.New(storage.DefaultOptions)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    err,
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    "Can't get body from request",
		})
		return
	}

	err = c.BindJSON(&link)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    "Can't parse JSON from body",
		})
		return
	}
	shortURL, err := generator.GenerateShortUrl(*link.URL, *link.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    err,
		})
		return
	}
	m.FullLink = *link.URL
	m.ShortLink = shortURL
	m.TTL = *link.TTL
	m.UserID = *link.UserID
	m.Count = 0
	err = redisDB.Put(shortURL, m)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   "OK",
		"msg":      "Link has been generated",
		"shotLink": "http://localhost:8080/" + shortURL,
	})

}

func RedirectByShortUrl(c *gin.Context) {
	redisCl := storage.ClientRedis{}
	redisDB, err := redisCl.New(storage.DefaultOptions)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    err,
		})
		return
	}
	shortUrl := c.Param("shortUrl")
	m, err := redisDB.Get(shortUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    err,
		})
		return
	}
	count := int64(m.Count)
	atomic.AddInt64(&count, 1)
	m.Count = uint(count)
	redisDB.Put(shortUrl, *m)
	c.Redirect(http.StatusTemporaryRedirect, m.FullLink)
	c.Abort()
}

func GetStatById(c *gin.Context) {
	redisCl := storage.ClientRedis{}
	redisDB, err := redisCl.New(storage.DefaultOptions)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    err,
		})
		return
	}
	shortUrl := c.Param("shortUrl")
	m, err := redisDB.Get(shortUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":         "OK",
		"msg":            "Stat info about current link",
		"shotLink":       "http://localhost:8080/" + shortUrl,
		"fullLink":       m.FullLink,
		"redirectsCount": m.Count,
		"userID":         m.UserID,
		"TTL":            m.TTL,
	})
}
