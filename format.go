package formatparser

const (
	TypeUnknown int = iota
	TypeText
	TypeKey
)

type FormatPart struct {
	Type  int
	Value string
}
