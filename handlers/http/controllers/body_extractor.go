package controllers

type BodyExtractor interface {
	Extract(url string) (host string, bodyStr string, responseTime int64, err error)
}
