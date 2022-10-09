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
	startTime := time.Now()
	log.Println("Web page analysis started")
	defer func(start time.Time) {
		log.Printf("Web page analysis completed. Time taken : %v ms", time.Since(start).Milliseconds())
	}(startTime)

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
	ginCtx.IndentedJSON(http.StatusOK, res)
	//
	//host, body, err := wpa.extractor.Extract(urlParam)
	//
	//if err != nil {
	//	log.Println(fmt.Sprintf("Error in getting the body for : %v", urlParam), err)
	//	ginCtx.IndentedJSON(
	//		http.StatusBadRequest,
	//		responses.NewErrorResponse("Error in getting the body for the given URL", err))
	//	return
	//}
	//
	//analyzerInfo := schema.NewAnalyzerInfo(body, host)
	//
	//wg.Add(len(wpa.analyzers))
	//for _, analyzer := range wpa.analyzers {
	//	go func(a analyzers.Analyzer, w *sync.WaitGroup) {
	//		a.Analyze(analyzerInfo, resManager)
	//		w.Done()
	//	}(analyzer, &wg)
	//}
	//wg.Wait()
	//log.Println("All analyzers completed")
	//ginCtx.IndentedJSON(http.StatusOK, resManager.To())
}

func (wpa *WebPageAnalyzerController) Analyze(resManager responses.WebPageAnalyzerResponseManager, url string) (
	response responses.AnalysisSuccessResponse, err error) {
	host, body, err := wpa.extractor.Extract(url)
	if err != nil {
		return responses.AnalysisSuccessResponse{}, err
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

	return resManager.To(), nil
}
