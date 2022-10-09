package analyzers

import (
	"github.com/naviud/webpage-analyzer/analyzers/schema"
	"github.com/naviud/webpage-analyzer/channels"
	"github.com/naviud/webpage-analyzer/handlers/http/responses"
	"golang.org/x/net/html"
	"log"
	"strings"
	"sync"
	"time"
)

type LinkProperty struct {
	Url        string
	Type       LinkType
	StatusCode int
	Latency    int64
}
type LinkType int

const (
	Internal LinkType = iota
	External
	linkHtmlTag = "a"
	href        = "href"
	http        = "http"
)

var wg sync.WaitGroup

type linkAnalyzer struct {
	links    sync.Map
	urlExec  channels.UrlExecutor
	provider channels.UrlExecutorProvider
}

func NewLinkAnalyzer(provider channels.UrlExecutorProvider) Analyzer {
	obj := linkAnalyzer{
		provider: provider,
	}
	return &obj
}

func (l *linkAnalyzer) Analyze(data schema.AnalyzerInfo, analysis responses.WebPageAnalyzerResponseManager) {
	startTime := time.Now()
	log.Println("Link analyzer started")
	defer func(start time.Time) {
		log.Printf("Link analyzer completed. Time taken : %v ms", time.Since(start).Milliseconds())
	}(startTime)

	l.links = sync.Map{}
	l.urlExec = l.provider.Provide()

	l.prepare(data)
	wg.Add(l.getMapLength())
	l.links.Range(func(key, value interface{}) bool {
		u := l.urlExec.Build(
			key.(string),
			&wg,
			func(key string, value int, latency int64) {
				linkProp, ok := l.links.Load(key)
				if !ok {
					log.Printf("Key : %v does not exist", key)
					return
				}
				tmpLinkProp := linkProp.(LinkProperty)
				tmpLinkProp.StatusCode = value
				tmpLinkProp.Latency = latency

				l.links.Store(key, tmpLinkProp)
			})
		u.PushChannel()
		return true
	})
	wg.Wait()
	l.setWebPageAnalyzer(analysis)
}

func (l *linkAnalyzer) prepare(data schema.AnalyzerInfo) {
	tokenizer := html.NewTokenizer(strings.NewReader(data.GetBody()))
	for {
		switch tokenizer.Next() {
		case html.StartTagToken:
			token := tokenizer.Token()
			if token.Data == linkHtmlTag {
				tmpUrl := GetTagAttribute(token, href)
				if isValidLink(tmpUrl) {
					tmpLinkOb := LinkProperty{
						Url:  tmpUrl,
						Type: getLinkType(tmpUrl, data.GetHost()),
					}
					l.links.Store(tmpUrl, tmpLinkOb)
				}
			}
		case html.ErrorToken:
			return
		}
	}
}

func (l *linkAnalyzer) getMapLength() int {
	var size int
	l.links.Range(func(key, value interface{}) bool {
		size++
		return true
	})
	return size
}

func (l *linkAnalyzer) setWebPageAnalyzer(analysis responses.WebPageAnalyzerResponseManager) {
	l.links.Range(func(key, value interface{}) bool {
		v := value.(LinkProperty)
		analysis.AddUrlInfo(v.Url, int(v.Type), v.StatusCode, v.Latency)
		return true
	})
}

func getLinkType(link string, host string) LinkType {
	if strings.Contains(link, host) {
		return Internal
	}
	return External
}

func isValidLink(link string) bool {
	trimmed := strings.TrimSpace(link)
	return strings.HasPrefix(trimmed, http)
}
