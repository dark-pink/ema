package file

// Content represents part of document.
type Content interface {
	// ContentType returns "data" or "text".
	// Value "data" represents type Data, value "text" represents type Text.
	ContentType() string

	content()
}
