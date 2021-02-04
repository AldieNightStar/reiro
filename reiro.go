package reiro

// Token representation as a result
type Token struct {
	Name string
	Data interface{}
}

// Parser function
type Parser func(p *Pointer) *Token

// NewToken - creates new Token
func NewToken(name string, data interface{}) *Token {
	return &Token{Name: name, Data: data}
}

// ParseToken - parses and get token. Will change Pointer position
// if no token found, it will save Pointer position as is
func ParseToken(p *Pointer, parsers []Parser) *Token {
	for _, parser := range parsers {
		pCopy := p.Copy()
		token := parser(pCopy)
		if token != nil {
			p.Apply(pCopy)
			return token
		}
	}

	return nil
}
