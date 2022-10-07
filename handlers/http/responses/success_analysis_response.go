package responses

import (
	"encoding/json"
	"log"
	"sync"
)

type AnalysisSuccessResponseManager struct {
	successRes AnalysisSuccessResponse
	lock       sync.RWMutex
}

type AnalysisSuccessResponse struct {
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

func NewWebPageAnalyzerResponseManager() *AnalysisSuccessResponseManager {
	headings := make([]Heading, 0)
	urls := make([]Url, 0)
	return &AnalysisSuccessResponseManager{
		successRes: AnalysisSuccessResponse{
			Headings: headings,
			Urls:     urls,
		},
	}
}

func (w *AnalysisSuccessResponseManager) SetHtmlVersion(htmlVersion string) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.successRes.HtmlVersion = htmlVersion
}

func (w *AnalysisSuccessResponseManager) SetTitle(title string) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.successRes.Title = title
}

func (w *AnalysisSuccessResponseManager) AddHeadingLevel(tag string, level string) {
	w.lock.Lock()
	defer w.lock.Unlock()

	fn := func() {
		levels := make([]string, 0)
		levels = append(levels, level)
		w.successRes.Headings = append(w.successRes.Headings, Heading{
			TagName: tag,
			Levels:  levels,
		})
	}

	if len(w.successRes.Headings) == 0 {
		fn()
	} else {
		for i, heading := range w.successRes.Headings {
			if heading.TagName == tag {
				heading.Levels = append(heading.Levels, level)
				w.successRes.Headings[i] = heading
				break
			}
			if i == len(w.successRes.Headings)-1 {
				fn()
			}
		}
	}
}

func (w *AnalysisSuccessResponseManager) AddUrlInfo(url string, urlType int, status int, latency int64) {
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
	w.successRes.Urls = append(w.successRes.Urls, u)
}

func (w *AnalysisSuccessResponseManager) SetHasLogin(hasLogin bool) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.successRes.HasLogin = hasLogin
}

func (w *AnalysisSuccessResponseManager) ToString() string {
	b, err := json.Marshal(w.successRes)
	if err != nil {
		log.Println("Error occurred when marshalling", err)
	}
	return string(b)
}

func (w *AnalysisSuccessResponseManager) To() AnalysisSuccessResponse {
	return w.successRes
}
