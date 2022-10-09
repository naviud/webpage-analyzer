package channels

import "sync"

type ExecFunc func(string, int, int64)

type UrlExecutor interface {
	Create() UrlExecutor
	Build(url string, wg *sync.WaitGroup, fn ExecFunc) UrlExecutor
	PushChannel()
}
