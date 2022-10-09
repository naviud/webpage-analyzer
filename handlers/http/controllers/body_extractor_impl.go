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

func (b *bodyExtractor) Extract(url string) (host string, bodyStr string, err error) {
	startTime := time.Now()
	log.Println("URL body fetching started")
	defer func(start time.Time) {
		log.Printf("URL body fetching completed. Time taken : %v ms", time.Since(start).Milliseconds())
	}(startTime)

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error occurred when getting the response", err)
		return "", "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error occurred when reading the response", err)
		return "", "", err
	}
	return resp.Request.Host, string(body), nil
}
