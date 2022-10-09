package channels

type UrlExecutorProvider interface {
	Provide() UrlExecutor
}
