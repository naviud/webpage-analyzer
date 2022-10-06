package channels

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"

	//"github.com/naviud/webpage-analyzer/analyzers"
	"log"
	"net/http"
	"sync"
)

type UrlExecutor struct {
	url string
	wg  *sync.WaitGroup
	fn  ExecFunc
}

type ExecFunc func(string, int, int64)

func NewUrlExecutor(url string, wg *sync.WaitGroup, fn ExecFunc) UrlExecutor {
	return UrlExecutor{
		url: url,
		wg:  wg,
		fn:  fn,
	}
}

var UrlExecutorChannel chan UrlExecutor

func InitUrlExecutorThreadPool() {
	UrlExecutorChannel = make(chan UrlExecutor)

	for i := 1; i <= 100; i++ {
		go executeUrl(UrlExecutorChannel, i)
	}
}

func executeUrl(channel chan UrlExecutor, id int) {
	log.Println(fmt.Sprintf("Thread stared : %d", id))

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
	//log.Println(fmt.Sprintf("Starting the request : %v", url))
	c := &http.Client{
		Timeout: 3 * time.Second,
	}
	res, err := c.Get(url)
	if err != nil {
		if err.(net.Error).Timeout() {
			fn(url, http.StatusGatewayTimeout, time.Since(startTime).Milliseconds())
		} else {
			fn(url, http.StatusBadRequest, time.Since(startTime).Milliseconds())
		}
		log.Println("error", url, err)
		return
	}
	defer res.Body.Close()
	fn(url, res.StatusCode, time.Since(startTime).Milliseconds())
}
