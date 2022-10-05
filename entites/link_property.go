package entites

type LinkType int

const (
	Internal LinkType = iota
	External
)

type LinkProperty struct {
	Url        string
	Type       LinkType
	StatusCode int
}
