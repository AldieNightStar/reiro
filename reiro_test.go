package reiro

import "testing"

func TestParseToken(t *testing.T) {
	parsers := []Parser{
		ParserVariableName,
		ParserNumber,
		ParserSpaces,
		ParserSymbols,
		ParserBrackets,
	}

	p := NewPointer("a\\1")
	if p.GetPointerValue() != 0 {
		t.Fatal("Pointer is not Zero Position")
	}

	t1 := ParseToken(p, parsers)

	if p.GetPointerValue() < 1 {
		t.Fatal("ParseToken is not incrementing Pointer")
	}

	t2 := ParseToken(p, parsers)

	if t1.Name == t2.Name {
		t.Fatal("Tokens should not be the same: " + t1.Name)
	}
}
