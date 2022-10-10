package channels

// UrlExecutorProvider interface contains the functions
// to be implemented to create the provider the builds
// UrlExecutor objects
type UrlExecutorProvider interface {

	// Provide function is responsible to provide
	// UrlExecutor objects
	Provide() UrlExecutor
}
