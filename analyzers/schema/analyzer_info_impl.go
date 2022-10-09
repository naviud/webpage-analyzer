package schema

type analyzerInfo struct {
	body string
	host string
}

func NewAnalyzerInfo(body string, host string) AnalyzerInfo {
	return &analyzerInfo{
		body: body,
		host: host,
	}
}

func (a *analyzerInfo) GetBody() string {
	return a.body
}

func (a *analyzerInfo) GetHost() string {
	return a.host
}
