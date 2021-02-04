package reiro

import "testing"

func TestPointerReadsChar(t *testing.T) {
	p := NewPointer("abc!")
	validateReadChar(t, p, 'a')
	validateReadChar(t, p, 'b')
	validateReadChar(t, p, 'c')
	validateReadChar(t, p, '!')
	validateReadChar(t, p, 0)
	validateReadChar(t, p, 0)
	validateReadChar(t, p, 0)
}

func TestPointerReadsCharAfterSeek(t *testing.T) {
	p := NewPointer("xyz!")
	validateReadChar(t, p, 'x')
	validateReadChar(t, p, 'y')
	p.SetPointerValue(0)
	validateReadChar(t, p, 'x')
	validateReadChar(t, p, 'y')
	validateReadChar(t, p, 'z')
	validateReadChar(t, p, '!')
	validateReadChar(t, p, 0)
}

func validateReadChar(t *testing.T, p *Pointer, r rune) {
	c := p.ReadChar()
	if c != r {
		t.Fatalf("Read char is not ok: [%c] but should be [%c]", r, c)
	}
	p.IncrementPointer()
}

func TestReadNextSyms(t *testing.T) {
	p := NewPointer("axyz")

	validateReadNextSym(t, p, 1, "a")
	validateReadNextSym(t, p, 2, "ax")
	validateReadNextSym(t, p, 3, "axy")
	validateReadNextSym(t, p, 4, "axyz")
	validateReadNextSym(t, p, 5, "axyz")
	validateReadNextSym(t, p, 20, "axyz")

	p.SetPointerValue(p.GetPointerValue() + 3)
	validateReadNextSym(t, p, 1, "z")
	validateReadNextSym(t, p, 5, "z")
}

func validateReadNextSym(t *testing.T, p *Pointer, n uint32, str string) {
	if p.ReadNextSyms(n) != str {
		t.Fatalf("Error reading next symbols: %s", str)
	}
}

func TestReadNextCheck(t *testing.T) {
	p := NewPointer("ax")
	if !p.IsNextText("ax") {
		t.Fatal("Check next symbols is not right")
	}
	if p.IsNextText("rr") {
		t.Fatal("Check next symbols is not right")
	}
}
