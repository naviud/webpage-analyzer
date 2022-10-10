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

	// Initialize and load configurations.
	configurations.Configurations{
		configurations.GetAppConfig(),
	}.Init()

	// Initialize URL executor thread pool.
	channels.InitUrlExecutorThreadPool(configurations.GetAppConfig().ChannelCount)

	// Initialize the URL body extractor.
	bodyExtractor := controllers.NewBodyExtractor()

	// Initialize the analyzers.
	htmlVersionAnalyzer := analyzers.NewHtmlVersionAnalyzer()
	titleAnalyzer := analyzers.NewTitleAnalyzer()
	headingAnalyzer := analyzers.NewHeadingAnalyzer()
	linkAnalyzer := analyzers.NewLinkAnalyzer(channels.NewUrlExecutorProvider(channels.NewUrlExecutor()))
	loginAnalyzer := analyzers.NewLoginFormAnalyzer()

	// Initialize the controller.
	controller := controllers.NewWebPageAnalyzerController(
		bodyExtractor,
		htmlVersionAnalyzer, titleAnalyzer, headingAnalyzer, linkAnalyzer, loginAnalyzer)

	// Initialize the HTTP server.
	svrDefault = http.Server{
		Addr:         fmt.Sprintf(":%v", configurations.GetAppConfig().ServicePort),
		Handler:      engines.NewDefaultEngine(controller).GetDefaultEngine(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	// Run the HTTP server.
	go func() {
		if err := svrDefault.ListenAndServe(); err != nil {
			log.Fatal("Failed to start the server", err)
		}
	}()

	<-channels.ChanSvrStart
	log.Printf("Service started under port 8080. To access the frontend, http://localhost:8080")

	<-sig
	log.Println("Shutting down...")
	if err := svrDefault.Shutdown(context.Background()); err != nil {
		log.Fatal("Failed to stop the server", err)
	}
}
