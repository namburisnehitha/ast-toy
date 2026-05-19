package asttoy

type TokenType int

const (
	FUN = iota
	IDENTIFIER
	LPAREN
	RPAREN
	COMMA
	PLUS
	RETURN
	LBRACE
	RBRACE
	INT
	NUM
	EOF
	ILLEGAL
	STRING
	ERROR
	NIL
)

type Token struct {
	TokenType TokenType
	Lexeme    string
	Line      int
	Literal   any
}
