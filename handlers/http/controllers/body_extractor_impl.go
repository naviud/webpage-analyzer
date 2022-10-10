package controllers

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type bodyExtractor struct {
}

func NewBodyExtractor() BodyExtractor {
	return &bodyExtractor{}
}

// Extract function is responsible to extract the
// properties of the given URL.
func (b *bodyExtractor) Extract(url string) (host string, bodyStr string, responseTime int64, err error) {
	startTime := time.Now()
	log.Println("URL body fetching started")

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error occurred when getting the response", err)
		return "", "", 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error occurred when reading the response", err)
		return "", "", 0, err
	}

	resTime := time.Since(startTime).Milliseconds()
	log.Printf("URL body fetching completed. Time taken : %v ms", resTime)

	return resp.Request.Host, string(body), resTime, nil
}
