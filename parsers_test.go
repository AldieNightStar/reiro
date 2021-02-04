package reiro

import "testing"

func TestParsers(t *testing.T) {
	txt := "abc 321.44 <f>##!"
	p := NewPointer(txt)

	parsers := []Parser{
		ParserVariableName,
		ParserNumber,
		ParserSpaces,
		ParserSymbols,
		ParserBrackets,
	}

	validateToken(t, p, parsers, "variableName", "abc")
	validateToken(t, p, parsers, "space", " ")
	validateToken(t, p, parsers, "number", "321.44")
	validateToken(t, p, parsers, "space", " ")
	validateToken(t, p, parsers, "brackets", "<")
	validateToken(t, p, parsers, "variableName", "f")
	validateToken(t, p, parsers, "brackets", ">")
	validateToken(t, p, parsers, "symbols", "##!")
}

func validateToken(t *testing.T, p *Pointer, parsers []Parser, name string, data string) {
	token := ParseToken(p, parsers)
	if token == nil {
		t.Fatal("Token is nil")
	}
	tokenDataStr, ok := token.Data.(string)
	if !ok {
		t.Fatal("Token data is not string")
	}
	if token.Name != name {
		t.Fatalf("Token is not [%s] as expected [%s]", token.Name, name)
	}
	if tokenDataStr != data {
		t.Fatalf("Token data [%s] is not as expected [%s]", tokenDataStr, data)
	}
}
