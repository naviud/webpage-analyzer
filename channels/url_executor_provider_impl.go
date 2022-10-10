package channels

type urlExecutorProvider struct {
	executor UrlExecutor
}

// NewUrlExecutorProvider function is responsible to create
// urlExecutorProvider objects
func NewUrlExecutorProvider(e UrlExecutor) UrlExecutorProvider {
	return &urlExecutorProvider{
		executor: e,
	}
}

// Provide function is responsible to provide
// UrlExecutor objects
func (u *urlExecutorProvider) Provide() UrlExecutor {
	return u.executor.Create()
}
