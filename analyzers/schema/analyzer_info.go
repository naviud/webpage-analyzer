package schema

type AnalyzerInfo interface {
	// GetBody function returns the body of the
	// web page.
	GetBody() string

	// GetHost function returns the host of the
	// web page.
	GetHost() string
}
