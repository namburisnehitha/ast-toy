package main

func Instrument(f *FuncDecl, info *FuncInfo) {

	selector := &SelectorExpr{
		X:   Ident{Name: "tracer"},
		Sel: Ident{Name: "Start"},
	}

	args := []Expr{
		&Ident{Name: info.ContextParamName},
		&StringLit{Value: f.Name.Name},
	}

	call := &CallExpr{
		Fun:  selector,
		Args: args,
	}

	ctx := Ident{Name: info.ContextParamName}
	span := Ident{Name: "span"}

	spanAssign := &AssignStmt{
		Lhs: []Ident{ctx, span},
		Rhs: []Expr{call},
	}

	selectordef := &SelectorExpr{
		X:   Ident{Name: "span"},
		Sel: Ident{Name: "End"},
	}

	calldef := &CallExpr{
		Fun:  selectordef,
		Args: nil,
	}

	spanDefer := &DeferStmt{
		Call: calldef,
	}

	f.Body.List = append([]Stmt{spanAssign, spanDefer}, f.Body.List...)

}
