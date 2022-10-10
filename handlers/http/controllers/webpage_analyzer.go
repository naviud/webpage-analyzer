package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/naviud/webpage-analyzer/analyzers"
	"github.com/naviud/webpage-analyzer/analyzers/schema"
	"github.com/naviud/webpage-analyzer/handlers/http/responses"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

type WebPageAnalyzerController struct {
	analyzers []analyzers.Analyzer
	extractor BodyExtractor
}

var wg sync.WaitGroup

func NewWebPageAnalyzerController(bodyExtractor BodyExtractor, analyzers ...analyzers.Analyzer) *WebPageAnalyzerController {
	wpa := WebPageAnalyzerController{}
	wpa.analyzers = append(wpa.analyzers, analyzers...)
	wpa.extractor = bodyExtractor
	return &wpa
}

func (wpa *WebPageAnalyzerController) AnalyzeWebPage(ginCtx *gin.Context) {
	resManager := responses.NewWebPageAnalyzerResponseManager()

	urlParam := strings.TrimSpace(ginCtx.Query("url"))

	res, err := wpa.Analyze(resManager, urlParam)

	if err != nil {
		log.Println(fmt.Sprintf("Error in getting the body for : %v", urlParam), err)
		ginCtx.IndentedJSON(
			http.StatusBadRequest,
			responses.NewErrorResponse("Error in getting the body for the given URL", err))
		return
	}
	log.Println("All analyzers completed")
	log.Printf("Web page analysis completed. Time taken : %v ms", res.ServiceTime)
	ginCtx.IndentedJSON(http.StatusOK, res)
}

func (wpa *WebPageAnalyzerController) Analyze(resManager responses.WebPageAnalyzerResponseManager, url string) (response responses.AnalysisSuccessResponse, err error) {
	startTime := time.Now()

	host, body, extractTime, err := wpa.extractor.Extract(url)
	if err != nil {
		return responses.AnalysisSuccessResponse{}, err
	}
	resManager.SetExtractTime(extractTime)
	analyzerInfo := schema.NewAnalyzerInfo(body, host)

	wg.Add(len(wpa.analyzers))
	for _, analyzer := range wpa.analyzers {
		go func(a analyzers.Analyzer, w *sync.WaitGroup) {
			a.Analyze(analyzerInfo, resManager)
			w.Done()
		}(analyzer, &wg)
	}
	wg.Wait()

	resManager.SetServiceTime(time.Since(startTime).Milliseconds())
	return resManager.To(), nil
}
