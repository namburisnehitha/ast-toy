package asttoy

type Node interface {
	nodeType() string
}

type Stmt interface {
	Node
	stmtgNode()
}

type Expr interface {
	Node
	exprNode()
}
