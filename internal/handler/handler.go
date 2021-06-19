package handler

import (
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
)

type Link struct {
	Url *string `json:"url"`
}

func CheckHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

func GenerateNewLink(c *gin.Context) {
	var url Link
	err := c.BindJSON(&url)
	if err != nil {
		c.JSON(http.StatusBadRequest, "ERROR")
		return
	}
	fmt.Printf("URL to store: %+v", *url.Url)
	c.JSON(http.StatusOK, gin.H{
		"got_url": url.Url,
	})

}

func RedirectByShortUrl(c *gin.Context) {
	// shortUrl := c.Param("shortUrl")
	// itirialUrl := retriveUrl(shortUrl)
	// c.Redirect(302, initialUrl)
}

func GetStatById(c *gin.Context) {

}
