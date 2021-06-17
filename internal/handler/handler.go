package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

func GetNewLink(c *gin.Context) {

}

func RedirectByShortUrl(c *gin.Context) {
	// shortUrl := c.Param("shortUrl")
	// itirialUrl := retriveUrl(shortUrl)
	// c.Redirect(302, initialUrl)
}

func GetStatById(c *gin.Context) {

}
