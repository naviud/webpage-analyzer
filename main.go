package main

import (
	"fmt"
	"github.com/naviud/webpage-analyzer/analyzers"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
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

	httpVersionAnalyzer := analyzers.NewHttpVersionAnalyzer()
	httpVersionAnalyzer.Analyze(htmlStr)
	log.Println("HTTP version", httpVersionAnalyzer.Get())

	titleAnalyzer := analyzers.NewTitleAnalyzer()
	titleAnalyzer.Analyze(htmlStr)
	log.Println("Title", titleAnalyzer.Get())

	headingAnalyzer := analyzers.NewHeadingAnalyzer()
	headingAnalyzer.Analyze(htmlStr)
	log.Println(fmt.Sprintf("%+v", headingAnalyzer.Get()))
}
