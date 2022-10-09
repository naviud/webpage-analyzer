package responses

// WebPageAnalyzerResponseManager This interface is responsible for setting
// the required items for the response that it supposes to provide
type WebPageAnalyzerResponseManager interface {

	// SetHtmlVersion For setting the html version
	SetHtmlVersion(htmlVersion string)

	// SetTitle For setting the page title
	SetTitle(title string)

	// AddHeadingLevel For adding the headers as well append the levels.
	AddHeadingLevel(tag string, level string)

	// AddUrlInfo For adding the URL information
	AddUrlInfo(url string, urlType int, status int, latency int64)

	// SetHasLogin For setting the login is available or not
	SetHasLogin(hasLogin bool)

	// ToString For getting the string representation of the response
	ToString() string

	// To For generating the response object
	To() AnalysisSuccessResponse
}
