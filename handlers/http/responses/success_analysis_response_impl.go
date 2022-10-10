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

// NewWebPageAnalyzerResponseManager function is responsible to create an
// object of WebPageAnalyzerResponseManager.
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

// SetHtmlVersion function is responsible for setting the html
// version.
func (w *analysisSuccessResponseManager) SetHtmlVersion(htmlVersion string) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.successRes.HtmlVersion = htmlVersion
}

// SetTitle function is responsible for setting the page title.
func (w *analysisSuccessResponseManager) SetTitle(title string) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.successRes.Title = title
}

// SetServiceTime function is responsible for setting the service
// time.
func (w *analysisSuccessResponseManager) SetServiceTime(serviceTime int64) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.successRes.ServiceTime = serviceTime
}

// SetExtractTime function is responsible for setting the web page
// extracting time.
func (w *analysisSuccessResponseManager) SetExtractTime(extractTime int64) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.successRes.WebExtractTime = extractTime
}

// AddHeadingLevel function is responsible for adding the headers
// as well append the levels.
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

// AddUrlInfo function is responsible for adding the URL information.
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

// SetHasLogin function is responsible for setting the login is
//available or not.
func (w *analysisSuccessResponseManager) SetHasLogin(hasLogin bool) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.successRes.HasLogin = hasLogin
}

// ToString  function is responsible for getting the string
//representation of the response.
func (w *analysisSuccessResponseManager) ToString() string {
	b, err := json.Marshal(w.successRes)
	if err != nil {
		log.Println("Error occurred when marshalling the response", err)
	}
	return string(b)
}

// To function is responsible for generating the response object.
func (w *analysisSuccessResponseManager) To() AnalysisSuccessResponse {
	return w.successRes
}
