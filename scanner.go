package main

import "strconv"

type Scanner struct {
	Source  string
	Start   int
	Current int
	Line    int
	Tokens  []Token
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		Source:  source,
		Start:   0,
		Current: 0,
		Line:    1,
		Tokens:  []Token{},
	}
}

func (s *Scanner) Scan() []Token {
	for s.Current != len(s.Source) {
		s.Start = s.Current
		s.scanToken()
	}
	s.Tokens = append(s.Tokens, Token{
		TokenType: EOF,
	})
	return s.Tokens
}

func (s *Scanner) advance() byte {
	current := s.Source[s.Current]
	s.Current += 1
	return current
}

func (s *Scanner) makeToken(tokenType TokenType, literal any) Token {
	return Token{
		TokenType: tokenType,
		Lexeme:    s.Source[s.Start:s.Current],
		Line:      s.Line,
		Literal:   literal,
	}
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case '(':
		s.Tokens = append(s.Tokens, s.makeToken(LPAREN, nil))
	case ')':
		s.Tokens = append(s.Tokens, s.makeToken(RPAREN, nil))
	case '{':
		s.Tokens = append(s.Tokens, s.makeToken(LBRACE, nil))
	case '}':
		s.Tokens = append(s.Tokens, s.makeToken(RBRACE, nil))
	case ',':
		s.Tokens = append(s.Tokens, s.makeToken(COMMA, nil))
	case ':':
		if s.peek() == '=' {
			s.advance()
			s.Tokens = append(s.Tokens, s.makeToken(DEFINE, nil))
		}
	case '.':
		s.Tokens = append(s.Tokens, s.makeToken(DOT, nil))
	case '"':
		for s.peek() != '"' && !s.isAtEnd() {
			s.advance()
		}
		s.advance()
		s.Tokens = append(s.Tokens, s.makeToken(STRINGLIT, nil))
	case ' ':
		break
	case '\t':
		break
	case '\n':
		s.Line += 1
	default:
		if isAlpha(c) {
			s.identifier()
		} else if isDigit(c) {
			s.number()
		} else {
			s.Tokens = append(s.Tokens, s.makeToken(ILLEGAL, nil))
		}
	}
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAlphaNumeric(c byte) bool {
	return isAlpha(c) || isDigit(c)
}

var keyWords = map[string]TokenType{
	"fun":    FUN,
	"return": RETURN,
	"int":    INT,
	"string": STRING,
	"error":  ERROR,
	"nil":    NIL,
	"defer":  DEFER,
}

func (s *Scanner) isAtEnd() bool {
	return s.Current >= len(s.Source)
}

func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return 0
	}
	return s.Source[s.Current]
}

func (s *Scanner) identifier() {
	for isAlphaNumeric(s.peek()) {
		s.advance()
	}
	text := s.Source[s.Start:s.Current]
	tokenType, ok := keyWords[text]
	if !ok {
		tokenType = IDENTIFIER
	}
	s.Tokens = append(s.Tokens, s.makeToken(tokenType, nil))
}

func (s *Scanner) number() {
	for isDigit(s.peek()) {
		s.advance()
	}
	num, err := strconv.Atoi(s.Source[s.Start:s.Current])
	if err != nil {
		s.Tokens = append(s.Tokens, s.makeToken(ILLEGAL, nil))
		return
	}
	s.Tokens = append(s.Tokens, s.makeToken(NUM, num))

}
