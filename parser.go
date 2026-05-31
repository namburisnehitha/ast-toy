package asttoy

import "fmt"

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
	if p.isAtEnd() {
        return Token{TokenType: EOF}
    }
	return p.Tokens[p.Current]
}

func (p *Parser) check(tokenType TokenType) bool {
	if p.isAtEnd() {
        return false
    }
	return p.Tokens[p.Current].TokenType == tokenType
}

func (p *Parser) expect(tokenType TokenType) Token {
	if p.check(tokenType) {
		return p.advance()
	}
	panic(fmt.Sprintf("expected %v, got %v at pos %d", tokenType, p.Tokens[p.Current].TokenType, p.Current))
}

func (p *Parser) parseIdent() *Ident {
    tok := p.expect(IDENTIFIER)
    return &Ident{Name: tok.Lexeme}
}

func (p *Parser) parseFieldList() *FieldList {
	fields := &FieldList{}
	for !p.check(RPAREN) {
		field := &Field{}

		field.Name = p.parseIdent()
		field.Type = p.parseType()

		if p.check(COMMA) {
			p.advance()
		}
		fields.Fields = append(fields.Fields, *field)
	}
	return fields
}

func (p *Parser) parseType() *Ident {
    tok := p.peek()
    switch tok.TokenType {
    case  STRING, ERROR:
        p.advance()
        return &Ident{Name: tok.Lexeme}
    default:
        panic("expected type")
    }
}


func (p *Parser) parseBlockStmt() *BlockStmt {
	stmts := &BlockStmt{}
	p.expect(LBRACE)
	for !p.check(RBRACE) {
		stmt := p.parseStmt()
		stmts.List = append(stmts.List, stmt)
	}
	p.expect(RBRACE)
	return stmts
}

func (p *Parser) parseReturnStmt() *ReturnStmt{
	stmts := &ReturnStmt{}
	p.expect(RETURN)
	for !p.check(RBRACE) && !p.check(EOF) {
		if p.check(IDENTIFIER) {
    		tok := p.advance()
    		stmts.Results = append(stmts.Results, &Ident{Name: tok.Lexeme})
		} else if p.check(NIL) {
    		p.advance()
   			stmts.Results = append(stmts.Results, &NilLit{})
		}else if p.check(COMMA) {
			p.advance()
		}
	}
	return stmts
}

func (p *Parser) parseStmt() Stmt{
	if p.check(RETURN){
		return p.parseReturnStmt()
	}else if p.check(IDENTIFIER){
		return p.parseAssignStmt()
	}else{
		panic("Invalid Input")
	}
}

func (p *Parser) parseFunc() *FuncDecl {

    p.expect(FUN)

    name := p.parseIdent()
    p.expect(LPAREN)
    params := p.parseFieldList()
    p.expect(RPAREN)
    
    var returnType *FieldList
    if p.check(LPAREN) {
    	p.expect(LPAREN)
    	returnType = p.parseFieldList()
    		p.expect(RPAREN)
	}else {
   		field := Field{Type: p.parseType()}
    	returnType = &FieldList{Fields: []Field{field}}
	}
    
    body := p.parseBlockStmt()

    return &FuncDecl{
		Name: name,
		Parameters : params,
		ReturnType : returnType,
		Body : body,
	}
}
