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
