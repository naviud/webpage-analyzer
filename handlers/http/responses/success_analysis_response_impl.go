package responses

import (
	"encoding/json"
	"log"
	"sync"
)

type analysisSuccessResponseManager struct {
	successRes AnalysisSuccessResponse
	lock       sync.RWMutex
}

type AnalysisSuccessResponse struct {
	HtmlVersion    string    `json:"htmlVersion"`
	Title          string    `json:"title"`
	ServiceTime    int64     `json:"serviceTime"`
	WebExtractTime int64     `json:"webExtractTime"`
	Headings       []Heading `json:"headings"`
	Urls           []Url     `json:"urls"`
	HasLogin       bool      `json:"hasLogin"`
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

func NewWebPageAnalyzerResponseManager() WebPageAnalyzerResponseManager {
	headings := make([]Heading, 0)
	urls := make([]Url, 0)
	return &analysisSuccessResponseManager{
		successRes: AnalysisSuccessResponse{
			Headings: headings,
			Urls:     urls,
		},
	}
}

func (w *analysisSuccessResponseManager) SetHtmlVersion(htmlVersion string) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.successRes.HtmlVersion = htmlVersion
}

func (w *analysisSuccessResponseManager) SetTitle(title string) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.successRes.Title = title
}

func (w *analysisSuccessResponseManager) SetServiceTime(serviceTime int64) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.successRes.ServiceTime = serviceTime
}

func (w *analysisSuccessResponseManager) SetExtractTime(extractTime int64) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.successRes.WebExtractTime = extractTime
}

func (w *analysisSuccessResponseManager) AddHeadingLevel(tag string, level string) {
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

func (w *analysisSuccessResponseManager) AddUrlInfo(url string, urlType int, status int, latency int64) {
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

func (w *analysisSuccessResponseManager) SetHasLogin(hasLogin bool) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.successRes.HasLogin = hasLogin
}

func (w *analysisSuccessResponseManager) ToString() string {
	b, err := json.Marshal(w.successRes)
	if err != nil {
		log.Println("Error occurred when marshalling the response", err)
	}
	return string(b)
}

func (w *analysisSuccessResponseManager) To() AnalysisSuccessResponse {
	return w.successRes
}
