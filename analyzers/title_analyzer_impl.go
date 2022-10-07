package analyzers

import (
	"fmt"
	"github.com/naviud/webpage-analyzer/analyzers/schema"
	"github.com/naviud/webpage-analyzer/handlers/http/responses"
	"golang.org/x/net/html"
	"log"
	"strings"
	"time"
)

const titleTag = "title"

type titleAnalyzer struct {
}

func NewTitleAnalyzer() Analyzer {
	return &titleAnalyzer{}
}

func (t *titleAnalyzer) Analyze(data *schema.AnalyzerInfo, analysis *responses.AnalysisSuccessResponseManager) {
	startTime := time.Now()
	log.Println("Title analyzer started")
	defer func(start time.Time) {
		log.Println(fmt.Sprintf("Title analyzer completed. Time taken : %v ms", time.Since(startTime).Milliseconds()))
	}(startTime)

	tokenizer := html.NewTokenizer(strings.NewReader(data.GetBody()))
	for {
		switch tokenizer.Next() {
		case html.StartTagToken:
			token := tokenizer.Token()
			if token.Data == titleTag {
				tokenizer.Next()
				analysis.SetTitle(tokenizer.Token().Data)
				return
			}
		case html.ErrorToken:
			return
		}
	}
}
