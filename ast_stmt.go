package main

type BlockStmt struct {
	List []Stmt
}

func (bs *BlockStmt) nodeType() string {
	return "BlockStmt"
}

func (bs *BlockStmt) stmtNode() {}

type ReturnStmt struct {
	Results []Expr
}

func (rs *ReturnStmt) nodeType() string {
	return "ReturnStmt"
}

func (rs *ReturnStmt) stmtNode() {}

type AssignStmt struct {
	Lhs []Ident
	Rhs []Expr
}

func (as *AssignStmt) nodeType() string {
	return "AssignStmt"
}

func (as *AssignStmt) stmtNode() {}

type ExprStmt struct {
	X Expr
}

func (es *ExprStmt) nodeType() string {
	return "ExprStmt"
}

func (es *ExprStmt) stmtNode() {}

type DeferStmt struct {
	Call Expr
}

func (ds *DeferStmt) nodeType() string {
	return "DeferStmt"
}

func (ds *DeferStmt) stmtNode() {}
