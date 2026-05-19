package asttoy

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

func (s *Scanner) makeToken(tokenType TokenType) Token {
	return Token{
		TokenType: tokenType,
		Lexeme:    s.Source[s.Start:s.Current],
		Line:      s.Line,
		Literal:   nil,
	}
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case '(':
		s.Tokens = append(s.Tokens, s.makeToken(LPAREN))
	case ')':
		s.Tokens = append(s.Tokens, s.makeToken(RPAREN))
	case '{':
		s.Tokens = append(s.Tokens, s.makeToken(LBRACE))
	case '}':
		s.Tokens = append(s.Tokens, s.makeToken(RBRACE))
	case ',':
		s.Tokens = append(s.Tokens, s.makeToken(COMMA))
	case '+':
		s.Tokens = append(s.Tokens, s.makeToken(PLUS))
	case ' ':
		break
	case '\t':
		break
	case '\n':
		s.Line += 1
	default:
		s.Tokens = append(s.Tokens, s.makeToken(ILLEGAL))
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
	"func":   FUN,
	"return": RETURN,
	"int":    INT,
}

func (s *Scanner) isAtEnd() bool {
	return s.Current >= len(s.Source)
}
