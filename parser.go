package asttoy

type Parser struct {
	Tokens  []Token
	Current int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{
		Tokens:  tokens,
		Current: 0,
	}
}

func (p *Parser) isAtEnd() bool {
	return p.Current >= len(p.Tokens)
}

func (p *Parser) advance() Token {
	current := p.Tokens[p.Current]
	p.Current += 1
	return current
}

func (p *Parser) peek() Token {
	return p.Tokens[p.Current]
}

func (p *Parser) check(tokenType TokenType) bool {
	return p.Tokens[p.Current].TokenType == tokenType
}

func (p *Parser) expect(tokenType TokenType) Token {
	if p.check(tokenType) {
		return p.advance()
	}
	panic("expected token ...")
}

func (p *Parser) parseFunc() {

	if p.check(FUN) {
	}

}
