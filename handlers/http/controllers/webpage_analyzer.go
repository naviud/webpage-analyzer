package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/naviud/webpage-analyzer/analyzers"
	"github.com/naviud/webpage-analyzer/analyzers/schema"
	"github.com/naviud/webpage-analyzer/handlers/http/responses"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

type WebPageAnalyzerController struct {
	analyzers []analyzers.Analyzer
}

var wg sync.WaitGroup

func NewWebPageAnalyzerController(analyzers ...analyzers.Analyzer) *WebPageAnalyzerController {
	wpa := WebPageAnalyzerController{}
	for _, analyzer := range analyzers {
		wpa.analyzers = append(wpa.analyzers, analyzer)
	}
	return &wpa
}

func (wpa *WebPageAnalyzerController) AnalyzeWebPage(ginCtx *gin.Context) {
	startTime := time.Now()
	log.Println("Web page analysis started")
	defer func(start time.Time) {
		log.Println(fmt.Sprintf("Web page analysis completed. Time taken : %v ms", time.Since(startTime).Milliseconds()))
	}(startTime)

	resManager := responses.NewWebPageAnalyzerResponseManager()

	urlParam := strings.TrimSpace(ginCtx.Query("url"))

	host, body, err := getBodyForUrl(urlParam)
	if err != nil {
		log.Println(fmt.Sprintf("Error in getting the body for : %v", urlParam), err)
		ginCtx.IndentedJSON(
			http.StatusBadRequest,
			responses.NewErrorResponse("Error in getting the body for the given URL", err))
		return
	}

	analyzerInfo := schema.NewAnalyzerInfo(body, host)

	wg.Add(len(wpa.analyzers))
	for _, analyzer := range wpa.analyzers {
		go func(a analyzers.Analyzer, w *sync.WaitGroup) {
			a.Analyze(analyzerInfo, resManager)
			w.Done()
		}(analyzer, &wg)
	}
	wg.Wait()
	log.Println("All analyzers completed")
	ginCtx.IndentedJSON(http.StatusOK, resManager.To())
}

func getBodyForUrl(url string) (host string, bodyStr string, err error) {
	startTime := time.Now()
	log.Println("URL body fetching started")
	defer func(start time.Time) {
		log.Println(fmt.Sprintf("URL body fetching completed. Time taken : %v ms", time.Since(startTime).Milliseconds()))
	}(startTime)

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error occurred when getting the response", err)
		return "", "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error occurred when reading the response", err)
		return "", "", err
	}
	return resp.Request.Host, string(body), nil
}
