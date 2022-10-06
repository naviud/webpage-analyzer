package analyzers

import (
	"fmt"
	"github.com/naviud/webpage-analyzer/analyzers/schema"
	"github.com/naviud/webpage-analyzer/handlers/http/responses"
	"golang.org/x/net/html"
	"log"
	"regexp"
	"strings"
	"time"
)

const headingHtmlTag = "[hH][1-9]"

type headingAnalyzer struct {
}

func NewHeadingAnalyzer() Analyzer {
	return &headingAnalyzer{}
}

func (h *headingAnalyzer) Analyze(data *schema.AnalyzerInfo, analysis *responses.WebPageAnalyzerResponseManager) {
	startTime := time.Now()
	log.Println("heading analyzer started")
	defer func(start time.Time) {
		log.Println(fmt.Sprintf("heading analyzer completed. Time taken : %v ms", time.Since(startTime).Milliseconds()))
	}(startTime)

	tokenizer := html.NewTokenizer(strings.NewReader(data.GetBody()))
	for {
	InnerLoopBreakLabel:
		switch tokenizer.Next() {
		case html.StartTagToken:
			token := tokenizer.Token()
			match, err := regexp.MatchString(headingHtmlTag, token.Data)
			if err != nil {
				log.Println("error in matching headings", err)
			}
			if match {
				tokenizer.Next()
				tmpTkn := tokenizer.Token()
				if tmpTkn.Type == html.TextToken {
					analysis.AddHeadingLevel(token.Data, tmpTkn.Data)
				} else {
					for {
						switch tokenizer.Next() {
						case html.TextToken:
							analysis.AddHeadingLevel(token.Data, tokenizer.Token().Data)
							break InnerLoopBreakLabel
						}
					}
				}
			}
		case html.ErrorToken:
			return
		}
	}
}
