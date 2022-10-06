package analyzers

import "golang.org/x/net/html"

func GetTagAttribute(token html.Token, name string) string {
	for _, attribute := range token.Attr {
		if attribute.Key == name {
			return attribute.Val
		}
	}
	return ""
}
