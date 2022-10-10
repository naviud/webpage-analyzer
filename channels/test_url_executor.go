package channels

import "sync"

// This is an implementation for the UrlExecutor interface
// in sake of unit testing purposes by mocking the actual
// behavior.
type testUrlExecutor struct {
	url string
	wg  *sync.WaitGroup
	fn  ExecFunc
}

func NewTestUrlExecutor() UrlExecutor {
	return &testUrlExecutor{}
}

func (u *testUrlExecutor) Create() UrlExecutor {
	return &testUrlExecutor{}
}

func (u *testUrlExecutor) Build(url string, wg *sync.WaitGroup, fn ExecFunc) UrlExecutor {
	u.url = url
	u.wg = wg
	u.fn = fn
	return u
}

func (u *testUrlExecutor) PushChannel() {
	u.wg.Done()
}
