package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/1r0npipe/shortener-web-links/internal/config"
	"github.com/1r0npipe/shortener-web-links/internal/params"
	"github.com/gin-gonic/gin"
)

func main() {
	flags, err := params.Init()
	if err != nil {
		fmt.Println(params.HelpMessage)
		log.Fatal("Can't read the flags properly")
	}
	fmt.Println(flags)
	config, err := config.ReadNewConfig(flags.FileConfig)
	if err != nil {
		log.Fatal("Can't read config file")
	}
	osSig := make(chan os.Signal, 1)
	signal.Notify(osSig,
		syscall.SIGINT,
		syscall.SIGTERM)

	r := gin.Default()
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	go r.Run(":" + config.Server.Port)

	<-osSig
	log.Printf("shutting down ...\n")
	_, cancel := context.WithTimeout(context.Background(), time.Duration(config.Server.Timeout)*time.Second)
	defer cancel()
	log.Printf("the service has been down...")

}
