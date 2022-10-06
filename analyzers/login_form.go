package analyzers

import (
	"github.com/naviud/webpage-analyzer/analyzers/schema"
	"github.com/naviud/webpage-analyzer/handlers/http/responses"
	"golang.org/x/net/html"
	"strings"
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

func (l *loginFormAnalyzer) Analyze(data *schema.AnalyzerInfo, analysis *responses.WebPageAnalyzerResponseManager) {
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
