package file

// Text represents part of document.
// It can be stored as plain text.
type Text []string

// ContentType always returns "text".
func (Text) ContentType() string {
	return "text"
}

func (Text) content() {
}

var _ Content = (Text)(nil)
