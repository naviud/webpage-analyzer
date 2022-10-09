package analyzers

import (
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

func (h *headingAnalyzer) Analyze(data schema.AnalyzerInfo, analysis responses.WebPageAnalyzerResponseManager) {
	startTime := time.Now()
	log.Println("Heading analyzer started")
	defer func(start time.Time) {
		log.Printf("Heading analyzer completed. Time taken : %v ms", time.Since(start).Milliseconds())
	}(startTime)

	regex, err := regexp.Compile(headingHtmlTag)
	if err != nil {
		log.Println("Error in compiling the regex", err)
		return
	}

	tokenizer := html.NewTokenizer(strings.NewReader(data.GetBody()))
	for {
	InnerLoopBreakLabel:
		switch tokenizer.Next() {
		case html.StartTagToken:
			token := tokenizer.Token()
			match := regex.Match([]byte(token.Data))
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
