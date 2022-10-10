package analyzers

import (
	"github.com/naviud/webpage-analyzer/analyzers/schema"
	"github.com/naviud/webpage-analyzer/handlers/http/responses"
	"golang.org/x/net/html"
	"log"
	"strings"
	"time"
)

const (
	inputHtmlTag          = "input"
	typeHtmlAttribute     = "type"
	pwdValueHtmlAttribute = "password"
)

type loginFormAnalyzer struct {
}

func NewLoginFormAnalyzer() Analyzer {
	return &loginFormAnalyzer{}
}

// Analyze function in this implementation is for analyzing
// the login forms. Having said this is supposed to analyze
// login forms, it analyzes the <input> tags with type as
// 'password'. Such tag is found, this quits from the loop.
func (l *loginFormAnalyzer) Analyze(data schema.AnalyzerInfo, analysis responses.WebPageAnalyzerResponseManager) {
	startTime := time.Now()
	log.Println("Login form analyzer started")
	defer func(start time.Time) {
		log.Printf("Login form analyzer completed. Time taken : %v ms", time.Since(start).Milliseconds())
	}(startTime)

	tokenizer := html.NewTokenizer(strings.NewReader(data.GetBody()))
	for {
		switch tokenizer.Next() {
		case html.StartTagToken:
		case html.SelfClosingTagToken:
			token := tokenizer.Token()
			if token.Data == inputHtmlTag && GetTagAttribute(token, typeHtmlAttribute) == pwdValueHtmlAttribute {
				analysis.SetHasLogin(true)
				return
			}
		case html.ErrorToken:
			return
		}
	}
}
