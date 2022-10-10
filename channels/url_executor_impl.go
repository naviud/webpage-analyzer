package channels

import (
	"github.com/naviud/webpage-analyzer/configurations"
	"net"
	"os"
	"os/signal"
	"sync/atomic"
	"time"

	//"github.com/naviud/webpage-analyzer/analyzers"
	"log"
	"net/http"
	"sync"
)

type urlExecutor struct {
	url string
	wg  *sync.WaitGroup
	fn  ExecFunc
}

var urlExecutorChannel chan urlExecutor
var ChanSvrStart = make(chan bool, 1)
var atomicVar int32

func NewUrlExecutor() UrlExecutor {
	return &urlExecutor{}
}

// Create function is responsible to create an object
// of UrlExecutor
func (u *urlExecutor) Create() UrlExecutor {
	return &urlExecutor{}
}

// Build function is responsible to build an already
// initialized UrlExecutor object.
func (u *urlExecutor) Build(url string, wg *sync.WaitGroup, fn ExecFunc) UrlExecutor {
	u.url = url
	u.wg = wg
	u.fn = fn
	return u
}

// PushChannel function is responsible to push the built
// UrlExecutor object to the channel.
func (u *urlExecutor) PushChannel() {
	urlExecutorChannel <- *u
}

// InitUrlExecutorThreadPool function is responsible to
// create goroutines for the provided thread count.
func InitUrlExecutorThreadPool(threadCount int) {
	urlExecutorChannel = make(chan urlExecutor)

	for i := 1; i <= threadCount; i++ {
		go executeUrl(urlExecutorChannel, i, threadCount)
	}
}

// executeUrl function is responsible to get the `urlExecutor`
// object of the channel and proceed the designated work. This
// function gets executed in goroutines and connected via a
// channel to the outside thread.
func executeUrl(channel chan urlExecutor, id int, maxId int) {
	log.Printf("Thread stared : %d", id)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	atomic.AddInt32(&atomicVar, 1)

	if int(atomicVar) == maxId {
		ChanSvrStart <- true
	}

SignalBreakLabel:
	for {
		select {
		case call := <-channel:
			execute(call.url, call.wg, call.fn)
		case <-signals:
			break SignalBreakLabel
		}
	}
}

// execute function calls the URL to make sure it is
// accessible which provides and executes the function
// which provides by passing the collected values from
// the URL execution.
func execute(url string, wg *sync.WaitGroup, fn func(url string, status int, latency int64)) {
	startTime := time.Now()
	defer wg.Done()
	c := &http.Client{
		Timeout: configurations.GetAppConfig().LinkTimeoutInMs,
	}
	res, err := c.Get(url)
	if err != nil {
		if err.(net.Error).Timeout() {
			fn(url, http.StatusGatewayTimeout, time.Since(startTime).Milliseconds())
		} else {
			fn(url, http.StatusBadRequest, time.Since(startTime).Milliseconds())
		}
		log.Println("Error in getting response", url, err)
		return
	}
	defer res.Body.Close()
	fn(url, res.StatusCode, time.Since(startTime).Milliseconds())
}
