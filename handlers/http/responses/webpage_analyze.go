package responses

import (
	"encoding/json"
	"log"
	"sync"
)

type WebPageAnalyzerResponseManager struct {
	webPageAnalyzerRes WebPageAnalyzerResponse
	lock               sync.RWMutex
}

type WebPageAnalyzerResponse struct {
	HtmlVersion string    `json:"htmlVersion"`
	Title       string    `json:"title"`
	Headings    []Heading `json:"headings"`
	Urls        []Url     `json:"urls"`
	HasLogin    bool      `json:"hasLogin"`
}

type Heading struct {
	TagName string   `json:"tagName"`
	Levels  []string `json:"levels"`
}

type Url struct {
	Url     string `json:"url"`
	Type    string `json:"type"`
	Status  int    `json:"status"`
	Latency int64  `json:"latency"`
}

func NewWebPageAnalyzerResponseManager() *WebPageAnalyzerResponseManager {
	headings := make([]Heading, 0)
	urls := make([]Url, 0)
	return &WebPageAnalyzerResponseManager{
		webPageAnalyzerRes: WebPageAnalyzerResponse{
			Headings: headings,
			Urls:     urls,
		},
	}
}

func (w *WebPageAnalyzerResponseManager) SetHtmlVersion(htmlVersion string) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.webPageAnalyzerRes.HtmlVersion = htmlVersion
}

func (w *WebPageAnalyzerResponseManager) SetTitle(title string) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.webPageAnalyzerRes.Title = title
}

func (w *WebPageAnalyzerResponseManager) AddHeadingLevel(tag string, level string) {
	w.lock.Lock()
	defer w.lock.Unlock()

	fn := func() {
		levels := make([]string, 0)
		levels = append(levels, level)
		w.webPageAnalyzerRes.Headings = append(w.webPageAnalyzerRes.Headings, Heading{
			TagName: tag,
			Levels:  levels,
		})
	}

	if len(w.webPageAnalyzerRes.Headings) == 0 {
		fn()
	} else {
		for i, heading := range w.webPageAnalyzerRes.Headings {
			if heading.TagName == tag {
				heading.Levels = append(heading.Levels, level)
				w.webPageAnalyzerRes.Headings[i] = heading
				break
			}
			if i == len(w.webPageAnalyzerRes.Headings)-1 {
				fn()
			}
		}
	}
}

func (w *WebPageAnalyzerResponseManager) AddUrlInfo(url string, urlType int, status int, latency int64) {
	urlTypeStr := "External"

	w.lock.Lock()
	defer w.lock.Unlock()

	if urlType == 0 {
		urlTypeStr = "Internal"
	}

	u := Url{
		Url:     url,
		Type:    urlTypeStr,
		Status:  status,
		Latency: latency,
	}
	w.webPageAnalyzerRes.Urls = append(w.webPageAnalyzerRes.Urls, u)
}

func (w *WebPageAnalyzerResponseManager) SetHasLogin(hasLogin bool) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.webPageAnalyzerRes.HasLogin = hasLogin
}

func (w *WebPageAnalyzerResponseManager) ToString() string {
	b, err := json.Marshal(w.webPageAnalyzerRes)
	if err != nil {
		log.Println("error occurred when marshalling", err)
	}
	return string(b)
}

func (w *WebPageAnalyzerResponseManager) To() WebPageAnalyzerResponse {
	return w.webPageAnalyzerRes
}
