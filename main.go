package main

import (
	"fmt"
	"github.com/naviud/webpage-analyzer/analyzers"
	"github.com/naviud/webpage-analyzer/channels"
	"github.com/naviud/webpage-analyzer/entites"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func main() {

	channels.InitUrlExecutorThreadPool()

	url := "https://www.red-gate.com/simple-talk/devops/testing/go-unit-tests-tips-from-the-trenches/"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	htmlStr := string(body)

	//httpVersionAnalyzer := analyzers.NewHttpVersionAnalyzer()
	//httpVersionAnalyzer.Analyze(htmlStr)
	//log.Println("HTTP version", httpVersionAnalyzer.Get())
	//
	//titleAnalyzer := analyzers.NewTitleAnalyzer()
	//titleAnalyzer.Analyze(htmlStr)
	//log.Println("Title", titleAnalyzer.Get())
	//
	//headingAnalyzer := analyzers.NewHeadingAnalyzer()
	//headingAnalyzer.Analyze(htmlStr)
	//log.Println(fmt.Sprintf("%+v", headingAnalyzer.Get()))

	linkAnalyzer := analyzers.NewLinkAnalyzer()
	linkAnalyzer.SetProperty(url)
	linkAnalyzer.Analyze(htmlStr)
	m := linkAnalyzer.Get().(sync.Map)

	m.Range(func(key, value interface{}) bool {
		log.Println(fmt.Sprintf("%+v", value.(entites.LinkProperty)))
		return true
	})
}
