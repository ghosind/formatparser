package formatparser

import "testing"

func getType(code int) string {
	switch code {
	case TypeText:
		return "text"
	case TypeKey:
		return "key"
	default:
		return "unknown"
	}
}

func Test_Parse(t *testing.T) {
	format := "[%time%zone] [%level] \\% %message"

	parts := Parse(format)
	expects := []FormatPart{
		{Type: TypeText, Value: "["},
		{Type: TypeKey, Value: "time"},
		{Type: TypeKey, Value: "zone"},
		{Type: TypeText, Value: "] ["},
		{Type: TypeKey, Value: "level"},
		{Type: TypeText, Value: "] \\% "},
		{Type: TypeKey, Value: "message"},
	}

	if len(parts) != len(expects) {
		t.Fatalf("Expect got %d format parts, actual %d", len(expects), len(parts))
	}

	for i, part := range parts {
		expect := expects[i]
		if part.Type != expect.Type {
			t.Errorf("Expect got %s type, actual %s type", getType(expect.Type), getType(part.Type))
		}
		if part.Value != expect.Value {
			t.Errorf("Expect got \"%s\", actual \"%s\"", expect.Value, part.Value)
		}
	}
}
