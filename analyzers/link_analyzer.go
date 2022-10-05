package analyzers

import "github.com/naviud/webpage-analyzer/properties"

type LinkAnalyzer interface {
	Analyzer
	properties.Property
}
