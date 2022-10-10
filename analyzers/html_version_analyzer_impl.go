package analyzers

import (
	"github.com/naviud/webpage-analyzer/analyzers/schema"
	"github.com/naviud/webpage-analyzer/handlers/http/responses"
	"log"
	"strings"
	"time"
)

const unknownVersion string = "Unknown"

type htmlVersionAnalyzer struct {
	types map[string]string
}

func NewHtmlVersionAnalyzer() Analyzer {
	obj := htmlVersionAnalyzer{types: make(map[string]string)}
	obj.types["HTML 4.01 Strict"] = `"-//W3C//DTD HTML 4.01//EN"`
	obj.types["HTML 4.01 Transitional"] = `"-//W3C//DTD HTML 4.01 Transitional//EN"`
	obj.types["HTML 4.01 Frameset"] = `"-//W3C//DTD HTML 4.01 Frameset//EN"`
	obj.types["XHTML 1.0 Strict"] = `"-//W3C//DTD XHTML 1.0 Strict//EN"`
	obj.types["XHTML 1.0 Transitional"] = `"-//W3C//DTD XHTML 1.0 Transitional//EN"`
	obj.types["XHTML 1.0 Frameset"] = `"-//W3C//DTD XHTML 1.0 Frameset//EN"`
	obj.types["XHTML 1.1"] = `"-//W3C//DTD XHTML 1.1//EN"`
	obj.types["HTML 5"] = `<!DOCTYPE html>`

	return &obj
}

// Analyze function is this implementation for the analyzing the
// html version in the provided web page. This function just matches
// the string for the provided html versions.
func (h *htmlVersionAnalyzer) Analyze(data schema.AnalyzerInfo, analysis responses.WebPageAnalyzerResponseManager) {
	startTime := time.Now()
	log.Println("Html version analyzer started")
	defer func(start time.Time) {
		log.Printf("Html version analyzer completed. Time taken : %v ms", time.Since(start).Milliseconds())
	}(startTime)

	version := unknownVersion
	for name, value := range h.types {
		if strings.Contains(data.GetBody(), strings.ToLower(value)) ||
			strings.Contains(data.GetBody(), strings.ToUpper(value)) ||
			strings.Contains(data.GetBody(), value) {
			version = name
			break
		}
	}
	analysis.SetHtmlVersion(version)
}
