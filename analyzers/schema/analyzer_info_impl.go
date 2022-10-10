package schema

type analyzerInfo struct {
	body string
	host string
}

// NewAnalyzerInfo function constructs the object by getting
// the attributes.
func NewAnalyzerInfo(body string, host string) AnalyzerInfo {
	return &analyzerInfo{
		body: body,
		host: host,
	}
}

// GetBody function returns the body of the
// web page.
func (a *analyzerInfo) GetBody() string {
	return a.body
}

// GetHost function returns the host of the
// web page.
func (a *analyzerInfo) GetHost() string {
	return a.host
}
