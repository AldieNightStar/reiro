package reiro

import "strings"

const allowedVariableNameSymbols = "$_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const allowedSpaceSymbols = " \t\n\r"
const allowedNumbers = "0123456789"
const allowedSymbols = "!@#$%^&*()_+.,;:'\\/\""
const allowedBrackets = "[]{}<>()"

// ParserVariableName parser
func ParserVariableName(p *Pointer) *Token {
	sb := &strings.Builder{}
	for {
		c := p.ReadChar()
		if strings.Index(allowedVariableNameSymbols, string(c)) == -1 {
			break
		}
		sb.WriteRune(c)
		p.IncrementPointer()
	}
	if sb.Len() == 0 {
		return nil
	}
	return NewToken("variableName", sb.String())
}

// ParserSpaces parser
func ParserSpaces(p *Pointer) *Token {
	sb := &strings.Builder{}
	for {
		c := p.ReadChar()
		if strings.Index(allowedSpaceSymbols, string(c)) == -1 {
			break
		}
		sb.WriteRune(c)
		p.IncrementPointer()
	}
	if sb.Len() == 0 {
		return nil
	}
	return NewToken("space", sb.String())
}

// ParserSymbols parser
func ParserSymbols(p *Pointer) *Token {
	sb := &strings.Builder{}
	for {
		c := p.ReadChar()
		if strings.Index(allowedSymbols, string(c)) == -1 {
			break
		}
		sb.WriteRune(c)
		p.IncrementPointer()
	}
	if sb.Len() == 0 {
		return nil
	}
	return NewToken("symbols", sb.String())
}

// ParserBrackets parser
func ParserBrackets(p *Pointer) *Token {
	sb := &strings.Builder{}
	for {
		c := p.ReadChar()
		if strings.Index(allowedBrackets, string(c)) == -1 {
			break
		}
		sb.WriteRune(c)
		p.IncrementPointer()
	}
	if sb.Len() == 0 {
		return nil
	}
	return NewToken("brackets", sb.String())
}

// ParserNumber parser
func ParserNumber(p *Pointer) *Token {
	sb := &strings.Builder{}
	wasDot := false
	for {
		c := p.ReadChar()
		if c == '.' {
			if sb.Len() == 0 {
				break
			}
			if wasDot {
				break
			}
			sb.WriteRune(c)
			p.IncrementPointer()
			wasDot = true
			continue
		}
		if strings.Index(allowedNumbers, string(c)) == -1 {
			break
		}
		sb.WriteRune(c)
		p.IncrementPointer()
	}
	if sb.Len() == 0 {
		return nil
	}
	nstr := sb.String()
	if strings.HasPrefix(nstr, ".") || strings.HasSuffix(nstr, ".") {
		return nil
	}
	return NewToken("number", sb.String())
}
