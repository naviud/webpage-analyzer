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

	sig := make(chan os.Signal, 0)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

	configurations.Configurations{
		configurations.GetAppConfig(),
	}.Init()

	channels.InitUrlExecutorThreadPool(configurations.GetAppConfig().ChannelCount)

	controller := controllers.NewWebPageAnalyzerController(
		analyzers.NewHtmlVersionAnalyzer(),
		analyzers.NewTitleAnalyzer(),
		analyzers.NewHeadingAnalyzer(),
		analyzers.NewLinkAnalyzer(),
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

	select {
	case <-sig:
		log.Println("Shutting down...")
		if err := svrDefault.Shutdown(context.Background()); err != nil {
			log.Fatal("Failed to stop the server", err)
		}
	}

	//startTime := time.Now()
	//
	//defer func(start time.Time) {
	//	log.Println(fmt.Sprintf("Time taken : %d", time.Since(start).Milliseconds()))
	//}(startTime)
	//
	//channels.InitUrlExecutorThreadPool()
	//
	//url := "https://www.red-gate.com/simple-talk/devops/testing/go-unit-tests-tips-from-the-trenches/"
	//resp, err := http.Get(url)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer resp.Body.Close()
	//
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//analyzerInfo := schema.NewAnalyzerInfo(string(body), resp.Request.Host)
	//
	//resManager := responses.NewWebPageAnalyzerResponseManager()
	//
	//a := make([]analyzers.Analyzer, 0)
	//
	//htmlVersionAnalyzer := analyzers.NewHtmlVersionAnalyzer()
	//titleAnalyzer := analyzers.NewTitleAnalyzer()
	//headingAnalyzer := analyzers.NewHeadingAnalyzer()
	//linkAnalyzer := analyzers.NewLinkAnalyzer()
	//loginFormAnalyzer := analyzers.NewLoginFormAnalyzer()
	//
	//a = append(a, htmlVersionAnalyzer)
	//a = append(a, titleAnalyzer)
	//a = append(a, headingAnalyzer)
	//a = append(a, linkAnalyzer)
	//a = append(a, loginFormAnalyzer)
	//
	//for _, analyzer := range a {
	//	analyzer.Analyze(analyzerInfo, resManager)
	//}
	//
	//log.Println(resManager.ToString())
}
