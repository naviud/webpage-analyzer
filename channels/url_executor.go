package channels

import (
	"fmt"
	"github.com/naviud/webpage-analyzer/entites"
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
	url      string
	linkProp *sync.Map
	wg       *sync.WaitGroup
}

type UrlExecutorResponse struct {
	Url        string
	StatusCode int
}

func NewUrlExecutor(url string, linkProp *sync.Map, wg *sync.WaitGroup) UrlExecutor {
	return UrlExecutor{
		url:      url,
		linkProp: linkProp,
		wg:       wg,
	}
}

var UrlExecutorChannel chan UrlExecutor

func InitUrlExecutorThreadPool() {
	UrlExecutorChannel = make(chan UrlExecutor)

	for i := 1; i <= 10; i++ {
		go executeUrl(UrlExecutorChannel)
	}
}

func executeUrl(channel chan UrlExecutor) {
	log.Println("Thread stared")

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

SignalBreakLabel:
	for {
		select {
		case call := <-channel:
			execute(call.url, call.linkProp, call.wg)
		case <-signals:
			break SignalBreakLabel
		}
	}
}

func execute(url string, m *sync.Map, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Println(fmt.Sprintf("Starting the request : %v", url))

	c := &http.Client{
		Timeout: 500 * time.Millisecond,
	}
	res, err := c.Get(url)
	if err != nil {
		if err.(net.Error).Timeout() {
			update(m, url, 504)
		}
		log.Println("error", url, err)
		return
	}
	defer res.Body.Close()
	update(m, url, res.StatusCode)
}

func update(m *sync.Map, url string, statusCode int) {
	v, ok := m.Load(url)
	if !ok {

	}
	v1 := v.(entites.LinkProperty)
	v1.StatusCode = statusCode

	m.Store(url, v1)
	log.Println("stored", url)
}
