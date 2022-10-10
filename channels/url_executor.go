package channels

import "sync"

type ExecFunc func(string, int, int64)

// UrlExecutor interface contains the functions to be
// implemented to execute a URL in a multi-threaded
//approach.
type UrlExecutor interface {

	// Create function is responsible to create an object
	// of UrlExecutor
	Create() UrlExecutor

	// Build function is responsible to build an already
	// initialized UrlExecutor object.
	Build(url string, wg *sync.WaitGroup, fn ExecFunc) UrlExecutor

	// PushChannel function is responsible to push the built
	// UrlExecutor object to the channel.
	PushChannel()
}
