package analyzers

import (
	"fmt"
	"github.com/naviud/webpage-analyzer/channels"
	"github.com/naviud/webpage-analyzer/entites"
	"golang.org/x/net/html"
	"log"
	"net/url"
	"strings"
	"sync"
)

const (
	linkHtmlTag = "a"
	href        = "href"
	http        = "http"
)

var wg sync.WaitGroup

type linkAnalyzer struct {
	links     sync.Map
	domainUrl string
}

func NewLinkAnalyzer() LinkAnalyzer {
	obj := linkAnalyzer{}
	return &obj
}

func (l *linkAnalyzer) Analyze(data interface{}) {
	var size int
	l.prepare(data)
	l.links.Range(func(key, value interface{}) bool {
		size++
		return true
	})

	wg.Add(size)
	l.links.Range(func(key, value interface{}) bool {
		u := channels.NewUrlExecutor(key.(string), &l.links, &wg)
		channels.UrlExecutorChannel <- u
		return true
	})
	wg.Wait()
}

func (l *linkAnalyzer) prepare(data interface{}) {
	tokenizer := html.NewTokenizer(strings.NewReader(data.(string)))
	for {
		switch tokenizer.Next() {
		case html.StartTagToken:
			token := tokenizer.Token()
			if token.Data == linkHtmlTag {
				tmpUrl := getAttribute(token, href)
				if isValidLink(tmpUrl) {
					tmpLinkOb := entites.LinkProperty{
						Url:  tmpUrl,
						Type: l.getLinkType(tmpUrl),
					}
					l.links.Store(tmpUrl, tmpLinkOb)
					log.Println(fmt.Sprintf("added : %+v", tmpLinkOb))
				}
			}
		case html.ErrorToken:
			return
		}
	}
}

func (l *linkAnalyzer) Get() interface{} {
	return l.links
}

func (l *linkAnalyzer) SetProperty(value interface{}) {
	l.domainUrl = value.(string)
}

func (l *linkAnalyzer) getLinkType(link string) entites.LinkType {
	url, err := url.Parse(l.domainUrl)
	if err != nil {

	}
	if strings.Contains(link, url.Host) {
		return entites.Internal
	}
	return entites.External
}

func getAttribute(token html.Token, name string) string {
	for _, attribute := range token.Attr {
		if attribute.Key == name {
			return attribute.Val
		}
	}
	return ""
}

func isValidLink(link string) bool {
	trimmed := strings.TrimSpace(link)
	return strings.HasPrefix(trimmed, http)
}
