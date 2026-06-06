package main

func main() {
	source := `fun HttpGet(ctx context, url string) (string, error) {
        return response, nil
}`

	tokens := NewScanner(source).Scan()
	f := NewParser(tokens).ParseFunc()

	info := AnalyseFunc(f)

	if ShouldInstrument(info) {
		Instrument(f, info)
	}

	PrintFunc(f)
}
