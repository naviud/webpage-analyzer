package analyzers

import (
	"golang.org/x/net/html"
	"strings"
)

type titleAnalyzer struct {
	title string
}

func NewTitleAnalyzer() Analyzer {
	return &titleAnalyzer{}
}

func (t *titleAnalyzer) Analyze(data interface{}) {
	tokenizer := html.NewTokenizer(strings.NewReader(data.(string)))
	for {
		switch tokenizer.Next() {
		case html.StartTagToken:
			token := tokenizer.Token()
			if token.Data == "title" {
				tokenizer.Next()
				t.title = tokenizer.Token().Data
				return
			}
		}
	}
}

func (t *titleAnalyzer) Get() interface{} {
	return t.title
}
