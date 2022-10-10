package controllers

// BodyExtractor interface contains the functions
// to be implemented to extract the web page body.
type BodyExtractor interface {

	// Extract function is responsible to extract the
	// properties of the given URL.
	Extract(url string) (host string, bodyStr string, responseTime int64, err error)
}
