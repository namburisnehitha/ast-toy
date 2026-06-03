package main

type Ident struct {
	Name string
}

func (i *Ident) nodeType() string { return "Ident" }
func (i *Ident) exprNode()        {}

type NilLit struct{}

func (n *NilLit) nodeType() string { return "NilLit" }
func (n *NilLit) exprNode()        {}

type SelectorExpr struct {
	X   Ident
	Sel Ident
}

func (s *SelectorExpr) exprNode()        {}
func (s *SelectorExpr) nodeType() string { return "SelectorExpr" }

type CallExpr struct {
	Fun  Expr
	Args []Expr
}

func (c *CallExpr) exprNode()        {}
func (c *CallExpr) nodeType() string { return "CallExpr" }

type StringLit struct {
	Value string
}

func (sl *StringLit) exprNode()        {}
func (sl *StringLit) nodeType() string { return "StringLit" }
