package controllers

type BodyExtractor interface {
	Extract(url string) (host string, bodyStr string, err error)
}
