package main

import (
	"fmt"
	"strings"
)

func printExpr(e Expr) {
	switch ex := e.(type) {
	case *Ident:
		fmt.Printf("%s", ex.Name)
	case *StringLit:
		fmt.Printf("%s", ex.Value)
	case *NilLit:
		fmt.Printf("nil")
	case *SelectorExpr:
		fmt.Printf("%s.%s", ex.X.Name, ex.Sel.Name)
	case *CallExpr:
		printExpr(ex.Fun)
		fmt.Printf("(")
		for i, arg := range ex.Args {
			if i > 0 {
				fmt.Printf(", ")
			}
			printExpr(arg)
		}
		fmt.Printf(")")
	}
}

func printStmt(s Stmt, indent int) {
	prefix := strings.Repeat("\t", indent)
	switch st := s.(type) {
	case *AssignStmt:
		fmt.Printf("%s", prefix)
		for i, id := range st.Lhs {
			if i > 0 {
				fmt.Printf(", ")
			}
			fmt.Printf("%s", id.Name)
		}
		fmt.Printf(" := ")
		printExpr(st.Rhs[0])
		fmt.Printf("\n")
	case *DeferStmt:
		fmt.Printf("%sdefer ", prefix)
		printExpr(st.Call)
		fmt.Printf("\n")
	case *ExprStmt:
		fmt.Printf("%s", prefix)
		printExpr(st.X)
		fmt.Printf("\n")
	case *ReturnStmt:
		fmt.Printf("%sreturn ", prefix)
		for i, r := range st.Results {
			if i > 0 {
				fmt.Printf(", ")
			}
			printExpr(r)
		}
		fmt.Printf("\n")
	}
}

func PrintFunc(f *FuncDecl) {

	fmt.Printf("func %s(", f.Name.Name)

	for i, field := range f.Parameters.Fields {
		if i > 0 {
			fmt.Printf(", ")
		}
		fmt.Printf("%s %s", field.Name.Name, field.Type.Name)
	}

	fmt.Printf(")")

	if len(f.ReturnType.Fields) == 1 {
		fmt.Printf(" %s", f.ReturnType.Fields[0].Type.Name)
	} else {
		fmt.Printf(" (")
		for i, field := range f.ReturnType.Fields {
			if i > 0 {
				fmt.Printf(", ")
			}
			fmt.Printf("%s", field.Type.Name)
		}
		fmt.Printf(")")
	}

	fmt.Printf(" {\n")

	for _, stmt := range f.Body.List {
		printStmt(stmt, 1)
	}
	fmt.Printf("}\n")
}
