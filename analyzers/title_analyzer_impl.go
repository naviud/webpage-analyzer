package analyzers

import (
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

func (t *titleAnalyzer) Analyze(data schema.AnalyzerInfo, analysis responses.WebPageAnalyzerResponseManager) {
	startTime := time.Now()
	log.Println("Title analyzer started")
	defer func(start time.Time) {
		log.Printf("Title analyzer completed. Time taken : %v ms", time.Since(start).Milliseconds())
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
