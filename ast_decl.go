package asttoy

type FuncDecl struct {
	Name       *Ident
	Parameters *FieldList
	ReturnType *FieldList
	Body       *BlockStmt
}

type Field struct {
	Name Node
	Type Node
}

type FieldList struct {
	fieldList []Field
}

func (f *FuncDecl) nodeType() string  { return "FuncDecl" }
func (f *FieldList) nodeType() string { return "FieldList" }
func (f *Field) nodeType() string     { return "Field" }
