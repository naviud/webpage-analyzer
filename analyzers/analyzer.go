package analyzers

type Analyzer interface {
	Analyze(data interface{})
	Get() interface{}
}
