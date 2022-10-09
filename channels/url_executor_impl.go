package channels

import (
	"github.com/naviud/webpage-analyzer/configurations"
	"net"
	"os"
	"os/signal"
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

//func NewUrlExecutor(url string, wg *sync.WaitGroup, fn ExecFunc) UrlExecutor {
//	return &urlExecutor{
//		url: url,
//		wg:  wg,
//		fn:  fn,
//	}
//}

var urlExecutorChannel chan urlExecutor

func NewUrlExecutor() UrlExecutor {
	return &urlExecutor{}
}

func (u *urlExecutor) Create() UrlExecutor {
	return &urlExecutor{}
}

func (u *urlExecutor) Build(url string, wg *sync.WaitGroup, fn ExecFunc) UrlExecutor {
	u.url = url
	u.wg = wg
	u.fn = fn
	return u
}

func (u *urlExecutor) PushChannel() {
	urlExecutorChannel <- *u
}

func InitUrlExecutorThreadPool(threadCount int) {
	urlExecutorChannel = make(chan urlExecutor)

	for i := 1; i <= threadCount; i++ {
		go executeUrl(urlExecutorChannel, i)
	}
}

func executeUrl(channel chan urlExecutor, id int) {
	log.Printf("Thread stared : %d", id)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

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
