package analyzers

import (
	"golang.org/x/net/html"
	"log"
	"regexp"
	"strings"
)

const headingHtmlTag = "[hH][1-9]"

type headingAnalyzer struct {
	headings map[string][]string
}

func NewHeadingAnalyzer() Analyzer {
	obj := headingAnalyzer{
		headings: make(map[string][]string),
	}
	return &obj
}

func (h *headingAnalyzer) Analyze(data interface{}) {
	tokenizer := html.NewTokenizer(strings.NewReader(data.(string)))
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
					h.headings[token.Data] = append(h.headings[token.Data], tmpTkn.Data)
				} else {
					for {
						switch tokenizer.Next() {
						case html.TextToken:
							h.headings[token.Data] = append(h.headings[token.Data], tokenizer.Token().Data)
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

func (h *headingAnalyzer) Get() interface{} {
	return h.headings
}
