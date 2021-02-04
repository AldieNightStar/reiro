package reiro

// Pointer oriented string
type Pointer struct {
	ptr uint32
	str string
}

// NewPointer creates new Pointer oriented string
func NewPointer(text string) *Pointer {
	p := &Pointer{
		ptr: 0,
		str: text,
	}
	return p
}

// ReadChar reads a char (No ptr + 1)
func (p *Pointer) ReadChar() rune {
	if int(p.ptr) > len(p.str)-1 {
		return 0
	}
	return rune(p.str[p.ptr])
}

// GetPointerValue gets an ptr value
func (p *Pointer) GetPointerValue() uint32 {
	return p.ptr
}

// SetPointerValue sets a new value for ptr
func (p *Pointer) SetPointerValue(ptr uint32) {
	p.ptr = ptr
}

// IncrementPointer adds 1 to pointer (ptr + 1)
func (p *Pointer) IncrementPointer() uint32 {
	p.ptr++
	return p.ptr
}

// ReadNextSyms reads text with q number of symbols after Pointer position (No ptr + 1)
func (p *Pointer) ReadNextSyms(q uint32) string {
	lastID := len(p.str) - 1
	lastCurrentID := p.ptr + q
	if lastCurrentID > uint32(lastID) {
		if int(p.ptr) <= lastID {
			return p.str[p.ptr:]
		}
		return ""
	}
	return p.str[p.ptr:int(p.ptr+q)]
}

// ReadUntilEnd reads all text after Pointer pos until end (No ptr + 1)
func (p *Pointer) ReadUntilEnd() string {
	if int(p.ptr) >= len(p.str)-1 {
		return ""
	}
	return p.str[p.ptr:]
}

// Copy creates Pointer copy
func (p *Pointer) Copy() *Pointer {
	return &Pointer{
		p.ptr,
		p.str,
	}
}

// Apply changes from another Pointer
func (p *Pointer) Apply(p2 *Pointer) {
	p.ptr = p2.ptr
	p.str = p2.str
}

// IsNextText - checks that next text is equals to txt
func (p *Pointer) IsNextText(txt string) bool {
	return p.ReadNextSyms(uint32(len(txt))) == txt
}
