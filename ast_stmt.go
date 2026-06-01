package asttoy

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
	ValueIdent []Ident
	ExpeIdent  []Expr
}

func (as *AssignStmt) nodeType() string {
	return "AssignStmt"
}

func (as *AssignStmt) stmtNode() {}
