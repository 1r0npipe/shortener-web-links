package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Link struct {
	Url *string `json:"url"`
	Ttl *int    `json:"ttl,omitempty"`
}

func CheckHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

func GenerateNewLink(c *gin.Context) {
	var link *Link
	err := c.BindJSON(&link)
	if err != nil {
		c.JSON(http.StatusBadRequest, "ERROR")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"got_url": link.Url,
		"ttl":     link.Ttl,
	})

}

func RedirectByShortUrl(c *gin.Context) {
	// shortUrl := c.Param("shortUrl")
	// itirialUrl := retriveUrl(shortUrl)
	// c.Redirect(302, initialUrl)
}

func GetStatById(c *gin.Context) {

}
