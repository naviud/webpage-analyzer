package schema

type AnalyzerInfo struct {
	body string
	host string
}

func NewAnalyzerInfo(body string, host string) *AnalyzerInfo {
	return &AnalyzerInfo{
		body: body,
		host: host,
	}
}

func (a *AnalyzerInfo) GetBody() string {
	return a.body
}

func (a *AnalyzerInfo) GetHost() string {
	return a.host
}
