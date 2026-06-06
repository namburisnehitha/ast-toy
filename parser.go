package main

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

func (p *Parser) isType(t Token) bool {
	switch t.TokenType {
	case STRING, ERROR, INT, IDENTIFIER:
		return true
	default:
		return false
	}
}

func (p *Parser) parseFieldList() *FieldList {
	fields := &FieldList{}
	for !p.check(RPAREN) && !p.check(EOF) {

		field := &Field{}

		if p.check(IDENTIFIER) && p.isType(p.peekNext()) {
			field.Name = p.parseIdent()
		}

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
	if p.isType(tok) {
		p.advance()
		return &Ident{Name: tok.Lexeme}
	}
	panic(fmt.Sprintf("expected type, got %v", tok.TokenType))
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

func (p *Parser) parseReturnStmt() *ReturnStmt {
	stmts := &ReturnStmt{}
	p.expect(RETURN)
	for !p.check(RBRACE) && !p.check(EOF) {
		if p.check(IDENTIFIER) {
			tok := p.advance()
			stmts.Results = append(stmts.Results, &Ident{Name: tok.Lexeme})
		} else if p.check(NIL) {
			p.advance()
			stmts.Results = append(stmts.Results, &NilLit{})
		} else if p.check(COMMA) {
			p.advance()
		} else {
			break
		}
	}
	return stmts
}

func (p *Parser) parseCallExpr() Expr {
	x := p.parseIdent()
	p.expect(DOT)
	sel := p.parseIdent()
	selexp := &SelectorExpr{
		X:   *x,
		Sel: *sel,
	}
	p.expect(LPAREN)

	var args []Expr
	for !p.check(RPAREN) {
		arg := p.parseExpr()
		if p.check(COMMA) {
			p.advance()
		}
		args = append(args, arg)
	}

	p.expect(RPAREN)

	return &CallExpr{
		Fun:  selexp,
		Args: args,
	}
}

func (p *Parser) parseAssignStmt() Stmt {

	var lhs []Ident
	lhs = append(lhs, *p.parseIdent())
	for p.check(COMMA) {
		p.advance()
		lhs = append(lhs, *p.parseIdent())
	}
	p.expect(DEFINE)

	var exprs []Expr
	expr := p.parseExpr()
	exprs = append(exprs, expr)
	return &AssignStmt{
		Lhs: lhs,
		Rhs: exprs,
	}
}

func (p *Parser) parseStmt() Stmt {
	if p.check(RETURN) {
		return p.parseReturnStmt()
	} else if p.check(IDENTIFIER) && p.peekNext().TokenType == DOT {
		return &ExprStmt{X: p.parseCallExpr()}
	} else if p.check(IDENTIFIER) {
		return p.parseAssignStmt()
	} else {
		panic("Invalid Input")
	}
}

func (p *Parser) peekNext() Token {
	if p.Current+1 >= len(p.Tokens) {
		return Token{TokenType: EOF}
	}
	return p.Tokens[p.Current+1]
}

func (p *Parser) parseExpr() Expr {
	if p.check(IDENTIFIER) && p.peekNext().TokenType == DOT {
		return p.parseCallExpr()
	}
	return p.parseIdent()
}

func (p *Parser) ParseFunc() *FuncDecl {

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
	} else {
		field := Field{Type: p.parseType()}
		returnType = &FieldList{Fields: []Field{field}}
	}

	body := p.parseBlockStmt()

	return &FuncDecl{
		Name:       name,
		Parameters: params,
		ReturnType: returnType,
		Body:       body,
	}
}
