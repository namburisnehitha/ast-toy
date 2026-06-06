package main

type FuncInfo struct {
	HasContextParam  bool
	ContextParamName string
	ReturnErr        bool
	IsInstrumented   bool
}

func AnalyseFunc(f *FuncDecl) *FuncInfo {

	info := FuncInfo{}

	for _, field := range f.Parameters.Fields {
		if field.Type.Name == "context" {
			info.HasContextParam = true
			info.ContextParamName = field.Name.Name
		}
	}

	for _, field := range f.ReturnType.Fields {
		if field.Type.Name == "error" {
			info.ReturnErr = true
		}
	}

	for _, stmt := range f.Body.List {
		switch s := stmt.(type) {
		case *AssignStmt:
			c, ok := s.Rhs[0].(*CallExpr)
			if ok {
				sel, ok := c.Fun.(*SelectorExpr)
				if ok && sel.Sel.Name == "Start" {
					info.IsInstrumented = true
				}
			}
		}
	}

	return &info

}

func ShouldInstrument(f *FuncInfo) bool {
	return f.HasContextParam == true && f.IsInstrumented == false
}
