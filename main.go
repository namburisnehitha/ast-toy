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
	fmt.Printf("%+v\n", ast)
}
