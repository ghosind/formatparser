package formatparser

const (
	// TypeUnknown is an unknown type token
	TypeUnknown int = iota
	// TypeText is a text string
	TypeText
	// TypeKey is a key string
	TypeKey
)

// FormatPart is a token part.
type FormatPart struct {
	// Type is the part type, it should be TypeText or TypeKey.
	Type int
	// Value is the string value.
	Value string
}
