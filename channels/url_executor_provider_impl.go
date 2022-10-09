package channels

type urlExecutorProvider struct {
	executor UrlExecutor
}

func NewUrlExecutorProvider(e UrlExecutor) UrlExecutorProvider {
	return &urlExecutorProvider{
		executor: e,
	}
}

func (u *urlExecutorProvider) Provide() UrlExecutor {
	return u.executor.Create()
}
