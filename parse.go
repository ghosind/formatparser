package formatparser

import "bytes"

// isAlnum indicates a character is an alphabet or number.
func isAlnum(c byte) bool {
	return (c >= '0' && c <= '9') ||
		(c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z')
}

// Parse parses a format string to formatted token parts.
func Parse(format string) []*FormatPart {
	parts := make([]*FormatPart, 0)
	part := new(FormatPart)
	isKey := false
	isEscaped := false
	buffer := bytes.Buffer{}
	length := len(format)

	for i := 0; i < length; i++ {
		c := format[i]

		if isEscaped && c != '%' {
			isEscaped = false
		}

		if !isKey {
			if c == '%' && !isEscaped {
				isKey = true
				if part.Type != TypeUnknown {
					part.Value = buffer.String()
					buffer = bytes.Buffer{}
					parts = append(parts, part)
					part = new(FormatPart)
				}

				part.Type = TypeKey

				continue
			} else if part.Type == TypeUnknown {
				part.Type = TypeText
			}

			if c == '\\' {
				isEscaped = true
			}
		} else {
			if !isAlnum(c) {
				isKey = false
				part.Value = buffer.String()
				buffer = bytes.Buffer{}
				parts = append(parts, part)
				part = new(FormatPart)
				i--
				continue
			}
		}

		buffer.WriteByte(c)
	}

	part.Value = buffer.String()
	parts = append(parts, part)

	return parts
}
