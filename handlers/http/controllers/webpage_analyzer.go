package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/naviud/webpage-analyzer/analyzers"
	"github.com/naviud/webpage-analyzer/analyzers/schema"
	"github.com/naviud/webpage-analyzer/handlers/http/responses"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
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
	resManager := responses.NewWebPageAnalyzerResponseManager()

	urlParam := strings.TrimSpace(ginCtx.Query("url"))
	resp, err := http.Get(urlParam)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	analyzerInfo := schema.NewAnalyzerInfo(string(body), resp.Request.Host)

	//for _, analyzer := range wpa.analyzers {
	//	analyzer.Analyze(analyzerInfo, resManager)
	//}

	wg.Add(len(wpa.analyzers))
	for _, analyzer := range wpa.analyzers {
		go func(a analyzers.Analyzer, w *sync.WaitGroup) {
			a.Analyze(analyzerInfo, resManager)
			w.Done()
		}(analyzer, &wg)
	}
	wg.Wait()

	ginCtx.IndentedJSON(http.StatusOK, resManager.To())
}
