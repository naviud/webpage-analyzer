package main

import (
	"context"
	"fmt"
	"github.com/naviud/webpage-analyzer/analyzers"
	"github.com/naviud/webpage-analyzer/channels"
	"github.com/naviud/webpage-analyzer/configurations"
	"github.com/naviud/webpage-analyzer/handlers/http/controllers"
	"github.com/naviud/webpage-analyzer/handlers/http/engines"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var svrDefault http.Server

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)

	configurations.Configurations{
		configurations.GetAppConfig(),
	}.Init()

	channels.InitUrlExecutorThreadPool(configurations.GetAppConfig().ChannelCount)
	p := channels.NewUrlExecutorProvider(channels.NewUrlExecutor())

	controller := controllers.NewWebPageAnalyzerController(
		controllers.NewBodyExtractor(),
		analyzers.NewHtmlVersionAnalyzer(),
		analyzers.NewTitleAnalyzer(),
		analyzers.NewHeadingAnalyzer(),
		analyzers.NewLinkAnalyzer(p),
		analyzers.NewLoginFormAnalyzer())

	svrDefault = http.Server{
		Addr:         fmt.Sprintf(":%v", configurations.GetAppConfig().ServicePort),
		Handler:      engines.NewDefaultEngine(controller).GetDefaultEngine(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	go func() {
		if err := svrDefault.ListenAndServe(); err != nil {
			log.Fatal("Failed to start the server", err)
		}
	}()

	<-sig
	log.Println("Shutting down...")
	if err := svrDefault.Shutdown(context.Background()); err != nil {
		log.Fatal("Failed to stop the server", err)
	}
}
