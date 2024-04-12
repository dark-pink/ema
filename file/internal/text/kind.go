package text

type Kind int

const (
	KindSpace = Kind(iota + 1)
	KindPunctuation
	KindAlphaNum
)
