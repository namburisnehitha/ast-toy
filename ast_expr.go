package asttoy

type Ident struct {
	Name string
}

func (i *Ident) nodeType() string {
	return "Ident"
}

func (i *Ident) exprNode() {}

type BasicLit struct {
	Value string
}

func (bl *BasicLit) nodeType() string {
	return "Literal"
}

func (bl *BasicLit) exprNode() {}
