package main

import (
	"fmt"
)

func main() {
	source := `func HttpGet(url string) (string, error) {
    span := tracer.start(ctx, context)
    return response, nil
}`

	tokens := NewScanner(source).Scan()

	ast := NewParser(tokens).ParseFunc()

	for _, stmt := range ast.Body.List {
		switch s := stmt.(type) {
		case *AssignStmt:
			callExpr := s.Rhs[0].(*CallExpr)
			fmt.Printf("CallExpr.Fun: %+v\n", callExpr.Fun)
			fmt.Printf("CallExpr.Args: %+v\n", callExpr.Args)
			fmt.Printf("AssignStmt: %+v\n", s)
			fmt.Printf("Rhs[0]: %+v\n", s.Rhs[0])
		case *ReturnStmt:
			fmt.Printf("ReturnStmt: %+v\n", s)
		}
	}

}
