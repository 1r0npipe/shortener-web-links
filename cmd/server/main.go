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
	"github.com/1r0npipe/shortener-web-links/internal/storage"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logger,err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	slog := logger.Sugar()
	slog.Info("Initializin the configuration ...")

	flags, err := params.Init()
	if err != nil {
		fmt.Println(params.HelpMessage)
		log.Fatal("Can't read the flags properly")
	}
	fmt.Println(flags)
	configData, err := config.ReadNewConfig(flags.FileConfig)
	if err != nil {
		log.Fatal("Can't read config file")
	}
	// override the port if provided at the CLI
	if flags.Port != "" {
		configData.Server.Port = flags.Port
	}
	redisCl := storage.ClientRedis{}
	// imported but not used
	redisDB, err := redisCl.New(storage.DefaultOptions)

	if err != nil {
		slog.Error(err)
	}
	osSig := make(chan os.Signal, 1)
	signal.Notify(osSig,
		syscall.SIGINT,
		syscall.SIGTERM)

	//router := gin.Default()
	router := gin.New()

	router.GET("/healthz", handler.CheckHealth)
	router.POST("/v1/link", handler.GenerateNewLink)
	// TODO:
	router.GET("/v1/:shortUrl", handler.RedirectByShortUrl)
	router.GET("/v1/stat/:shortUrl", handler.GetStatById)
	// TODO:
	go func() {
		err := router.Run(":" + configData.Server.Port)
		if err != nil {
			log.Fatal("can't start server")
		}
	}()
	<-osSig
	log.Printf("shutting down ...\n")
	_, cancelFunc := context.WithTimeout(context.Background(), time.Duration(configData.Server.Timeout)*time.Second)
	defer cancelFunc()
	log.Printf("the service has been down...")

}
