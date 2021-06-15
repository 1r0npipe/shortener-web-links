package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/1r0npipe/shortener-web-links/internal/config"
	"github.com/1r0npipe/shortener-web-links/internal/handler"
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
	// override the port if provided at the CLI
	if flags.Port != "" {
		config.Server.Port = flags.Port
	}
	osSig := make(chan os.Signal, 1)
	signal.Notify(osSig,
		syscall.SIGINT,
		syscall.SIGTERM)

	//router := gin.Default()
	router := gin.New()

	router.GET("/healthz", handler.CheckHealth)
	// TODO:
	router.GET("/link", handler.GetNewLink)
	router.GET("/link/:{id}", handler.GetLinkById)
	router.GET("/stat/:{id}", handler.GetStatById)

	// TODO:
	go func() {
		err := router.Run(":" + config.Server.Port)
		if err != nil {
			log.Fatal("can't start server")
		}
	}()
	<-osSig
	log.Printf("shutting down ...\n")
	_, cancelFunc := context.WithTimeout(context.Background(), time.Duration(config.Server.Timeout)*time.Second)
	defer cancelFunc()
	log.Printf("the service has been down...")

}
