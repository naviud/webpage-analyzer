package responses

// WebPageAnalyzerResponseManager interface contains the functions
// to be implemented in setting the required items to the response
// that it supposes to provide.
type WebPageAnalyzerResponseManager interface {

	// SetHtmlVersion function is responsible for setting the html
	// version.
	SetHtmlVersion(htmlVersion string)

	// SetTitle function is responsible for setting the page title.
	SetTitle(title string)

	// SetServiceTime function is responsible for setting the service
	// time.
	SetServiceTime(serviceTime int64)

	// SetExtractTime function is responsible for setting the web page
	// extracting time.
	SetExtractTime(extractTime int64)

	// AddHeadingLevel function is responsible for adding the headers
	// as well append the levels.
	AddHeadingLevel(tag string, level string)

	// AddUrlInfo function is responsible for adding the URL information.
	AddUrlInfo(url string, urlType int, status int, latency int64)

	// SetHasLogin function is responsible for setting the login is
	//available or not.
	SetHasLogin(hasLogin bool)

	// ToString  function is responsible for getting the string
	//representation of the response.
	ToString() string

	// To function is responsible for generating the response object.
	To() AnalysisSuccessResponse
}
