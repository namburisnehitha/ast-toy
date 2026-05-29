package asttoy

type Ident struct {
	Name string
}

func (i *Ident) nodeType() string {return "Ident"}
func (i *Ident) exprNode() {}


type NilLit struct{}

func (n *NilLit) nodeType() string { return "NilLit" }
func (n *NilLit) exprNode()        {}
