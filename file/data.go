package file

// Data represents part of document.
// It can be stored as JSON object.
type Data map[string]any

// ContentType always returns "data".
func (Data) ContentType() string {
	return "data"
}

func (Data) content() {
}

var _ Content = (Data)(nil)
