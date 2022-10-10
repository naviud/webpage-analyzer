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
	a      = "a"
	href   = "href"
	http   = "http"
	script = "script"
	src    = "src"
	link   = "link"
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

// Analyze function this implementation is for analysing the
// links that the page contains. This analyzes the following
// combinations.
// <a href="">     => Hyper links
// <script src=""> => JS links
// <link href="">  => Links for external style sheets
// Any other required rules can be added to this.
//
// This function collects the links available in the page via the
// `prepare` function and such collected links(URLs) are sent to
// a thread pool to make sure the links are accessible or not in
// a parallel fashion. To pass these links to the thread pool,
// channels are used and in order to get a consolidated output,
// workgroups are used.
func (l *linkAnalyzer) Analyze(data schema.AnalyzerInfo, analysis responses.WebPageAnalyzerResponseManager) {
	startTime := time.Now()
	log.Println("Link analyzer started")
	defer func(start time.Time) {
		log.Printf("Link analyzer completed. Time taken : %v ms", time.Since(start).Milliseconds())
	}(startTime)

	// Initialize the sync map
	l.links = sync.Map{}

	// Initialize the UrlExecutor object
	l.urlExec = l.provider.Provide()

	// This adds the links to the map.
	l.prepare(data)

	// Create the worker with size of the
	// length of the map.
	wg.Add(l.getMapLength())

	// Iterate the map to build the object to be passed
	// to the channel.
	l.links.Range(func(key, value interface{}) bool {
		// Build the UrlExecutor object.
		u := l.urlExec.Build(
			key.(string),
			&wg,
			// The function that passes is responsible to
			// get execution details of the link (ex:
			// http status and latency) and update them to
			// the map against the key.
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
		// Push the UrlExecutor object to the channel.
		u.PushChannel()
		return true
	})
	// Wait until all the workers are done with their job.
	wg.Wait()
	l.setWebPageAnalyzer(analysis)
}

// prepare function is responsible to tokenize the web page and
// finds out the links in following rules.
// <a href="">     => Hyper links
// <script src=""> => JS links
// <link href="">  => Links for external style sheets
func (l *linkAnalyzer) prepare(data schema.AnalyzerInfo) {
	tokenizer := html.NewTokenizer(strings.NewReader(data.GetBody()))
	for {
		switch tokenizer.Next() {
		case html.StartTagToken:
			token := tokenizer.Token()
			if token.Data == a {
				tmpUrl := GetTagAttribute(token, href)

				//Only the valid links are added to the map.
				if isValidLink(tmpUrl) {
					tmpLinkOb := LinkProperty{
						Url:  tmpUrl,
						Type: getLinkType(tmpUrl, data.GetHost()),
					}
					l.links.Store(tmpUrl, tmpLinkOb)
				}
			}
			if token.Data == script {
				tmpUrl := GetTagAttribute(token, src)

				//Only the valid links are added to the map.
				if isValidLink(tmpUrl) {
					tmpLinkOb := LinkProperty{
						Url:  tmpUrl,
						Type: getLinkType(tmpUrl, data.GetHost()),
					}
					l.links.Store(tmpUrl, tmpLinkOb)
				}
			}
			if token.Data == link {
				tmpUrl := GetTagAttribute(token, href)

				//Only the valid links are added to the map.
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

// getMapLength function is responsible to get the length
// of the sync map that uses to store the captured links.
func (l *linkAnalyzer) getMapLength() int {
	var size int
	l.links.Range(func(key, value interface{}) bool {
		size++
		return true
	})
	return size
}

// setWebPageAnalyzer function is responsible to set the
// stored values in the sync map to the object of
// WebPageAnalyzerResponseManager.
func (l *linkAnalyzer) setWebPageAnalyzer(analysis responses.WebPageAnalyzerResponseManager) {
	l.links.Range(func(key, value interface{}) bool {
		v := value.(LinkProperty)
		analysis.AddUrlInfo(v.Url, int(v.Type), v.StatusCode, v.Latency)
		return true
	})
}

// getLinkType function is responsible to identify the
// nature of the link as internal or external. If link
// contains the host of the provided link, then that is
// internal. Otherwise, it's external.
func getLinkType(link string, host string) LinkType {
	if strings.Contains(link, host) {
		return Internal
	}
	return External
}

// isValidLink function is responsible to identify the
// link is valid or not. If the link contains the "http"
// phrase, such links are valid. Otherwise, it's invalid.
func isValidLink(link string) bool {
	trimmed := strings.TrimSpace(link)
	return strings.HasPrefix(trimmed, http)
}
