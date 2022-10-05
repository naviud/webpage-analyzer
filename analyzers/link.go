package analyzers

type LinkType int

const (
	Internal LinkType = iota
	External
)

type LinkProperty struct {
	Name       string
	Type       LinkType
	Accessible bool
}

type linkAnalyzer struct {
	links map[string]LinkProperty
}

func NewLinkAnalyzer() Analyzer {
	obj := linkAnalyzer{
		links: make(map[string]LinkProperty),
	}
	return &obj
}

func (l *linkAnalyzer) Analyze(data interface{}) {

}

func (l *linkAnalyzer) Get() interface{} {
	return l.links
}
