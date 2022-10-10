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

// NewWebPageAnalyzerController function is responsible to
// create WebPageAnalyzerController object by collecting the
// body extractor and analyzers.
func NewWebPageAnalyzerController(bodyExtractor BodyExtractor, analyzers ...analyzers.Analyzer) *WebPageAnalyzerController {
	wpa := WebPageAnalyzerController{}
	wpa.analyzers = append(wpa.analyzers, analyzers...)
	wpa.extractor = bodyExtractor
	return &wpa
}

// AnalyzeWebPage function is the controller for the
// `/v1/analyze` endpoint.
func (wpa *WebPageAnalyzerController) AnalyzeWebPage(ginCtx *gin.Context) {
	resManager := responses.NewWebPageAnalyzerResponseManager()

	// Get the passed URL from the query parameter.
	urlParam := strings.TrimSpace(ginCtx.Query("url"))

	// Refer to the comment of the function.
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

// Analyze function is responsible for executing the
// analyzers by providing the extracted body of the
// given URL. These analyzers are executed via goroutines
// within a work group for a parallelized execution.
func (wpa *WebPageAnalyzerController) Analyze(resManager responses.WebPageAnalyzerResponseManager, url string) (response responses.AnalysisSuccessResponse, err error) {
	startTime := time.Now()

	// Extract the body from the URL
	host, body, extractTime, err := wpa.extractor.Extract(url)
	if err != nil {
		return responses.AnalysisSuccessResponse{}, err
	}
	resManager.SetExtractTime(extractTime)
	analyzerInfo := schema.NewAnalyzerInfo(body, host)

	// Create work group with the size of analyzers
	wg.Add(len(wpa.analyzers))

	// Iterates the analyzers and create goroutine for each.
	for _, analyzer := range wpa.analyzers {

		// This function executes the `Analyze` function and
		// once that is done, mark it as done in work group.
		go func(a analyzers.Analyzer, w *sync.WaitGroup) {
			a.Analyze(analyzerInfo, resManager)
			w.Done()
		}(analyzer, &wg)
	}
	// Waits until all the goroutines are done with the executions.
	wg.Wait()

	resManager.SetServiceTime(time.Since(startTime).Milliseconds())
	return resManager.To(), nil
}
