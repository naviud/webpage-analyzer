package analyzers

import "strings"

const unknownVersion string = "Unknown"

type httpVersionAnalyzer struct {
	types   map[string]string
	version string
}

func NewHttpVersionAnalyzer() Analyzer {
	obj := httpVersionAnalyzer{types: make(map[string]string)}
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

func (h *httpVersionAnalyzer) Analyze(data interface{}) {
	h.version = unknownVersion
	strData := data.(string)
	for name, value := range h.types {
		if strings.Contains(strData, value) {
			h.version = name
			break
		}
	}
}

func (h *httpVersionAnalyzer) Get() interface{} {
	return h.version
}
