package analyzers

import (
	"github.com/naviud/webpage-analyzer/analyzers/schema"
	"github.com/naviud/webpage-analyzer/handlers/http/responses"
	"golang.org/x/net/html"
	"strings"
)

type titleAnalyzer struct {
}

func NewTitleAnalyzer() Analyzer {
	return &titleAnalyzer{}
}

func (t *titleAnalyzer) Analyze(data *schema.AnalyzerInfo, analysis *responses.WebPageAnalyzerResponseManager) {
	tokenizer := html.NewTokenizer(strings.NewReader(data.GetBody()))
	for {
		switch tokenizer.Next() {
		case html.StartTagToken:
			token := tokenizer.Token()
			if token.Data == "title" {
				tokenizer.Next()
				analysis.SetTitle(tokenizer.Token().Data)
				return
			}
		case html.ErrorToken:
			return
		}
	}
}
