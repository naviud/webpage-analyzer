package analyzers

import (
	"github.com/naviud/webpage-analyzer/analyzers/schema"
	"github.com/naviud/webpage-analyzer/handlers/http/responses"
)

// Analyzer interface contains the functions to be implemented
// to analyze web pages for given rules.
type Analyzer interface {

	// Analyze function analyzes the web page by the given
	// rules and outputs to the WebPageAnalyzerResponseManager.
	Analyze(schema.AnalyzerInfo, responses.WebPageAnalyzerResponseManager)
}
