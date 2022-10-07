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

func (l *loginFormAnalyzer) Analyze(data *schema.AnalyzerInfo, analysis *responses.AnalysisSuccessResponseManager) {
	startTime := time.Now()
	log.Println("Login form analyzer started")
	defer func(start time.Time) {
		log.Println(fmt.Sprintf("Login form analyzer completed. Time taken : %v ms", time.Since(startTime).Milliseconds()))
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
