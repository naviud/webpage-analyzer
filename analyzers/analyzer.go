package analyzers

import (
	"github.com/naviud/webpage-analyzer/analyzers/schema"
	"github.com/naviud/webpage-analyzer/handlers/http/responses"
)

type Analyzer interface {
	Analyze(*schema.AnalyzerInfo, *responses.WebPageAnalyzerResponseManager)
}
