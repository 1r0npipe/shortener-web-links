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

func GetLinkById(c *gin.Context) {

}

func GetStatById(c *gin.Context) {

}

func ForwardByLink(c *gin.Context) {

}
