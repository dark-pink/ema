package text

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"unicode"

	"darkp.ink/ema/file/internal"
)

func Parse(reader *internal.Reader) ([]string, error) {
	for sb, tokens, kind := (strings.Builder{}), []string(nil), Kind(0); ; {
		r, err := reader.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				tokens = append(tokens, sb.String())
				return tokens, nil
			}
			return nil, fmt.Errorf("text.Parse: %w", err)
		}
		if r == '{' {
			r, err = reader.ReadRune()
			if err != nil {
				return nil, fmt.Errorf("text.Parse: %w", err)
			}
			if r != '{' {
				reader.UnreadRune()
				reader.UnreadRune()
				tokens = append(tokens, sb.String())
				return tokens, nil
			}
		}
		switch {
		case unicode.IsSpace(r):
			if kind != KindSpace {
				if kind != 0 {
					tokens = append(tokens, sb.String())
					sb.Reset()
				}
				kind = KindSpace
			}
		case unicode.IsPunct(r):
			if kind != KindPunctuation {
				if kind != 0 {
					tokens = append(tokens, sb.String())
					sb.Reset()
				}
				kind = KindPunctuation
			}
		default:
			if kind != KindAlphaNum {
				if kind != 0 {
					tokens = append(tokens, sb.String())
					sb.Reset()
				}
				kind = KindAlphaNum
			}
		}
		sb.WriteRune(r)
	}
}
