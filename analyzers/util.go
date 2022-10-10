package analyzers

import "golang.org/x/net/html"

// GetTagAttribute function is responsible to get the
// value for the given tag name.
func GetTagAttribute(token html.Token, name string) string {
	for _, attribute := range token.Attr {
		if attribute.Key == name {
			return attribute.Val
		}
	}
	return ""
}
