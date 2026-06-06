# Ast-Toy

## What it is
It's a well scoped version of otel compile time instrument as notes for me to understand otelc better

## What it does 
Converts
```
func HttpGet(url string) (string, error) {
    return response, nil
}
```
To
```
func HttpGet(url string) (string, error) {
  ctx, span := tracer.Start(ctx,"HttpGet")
  defer span.End()
  return response, nil
}
```

## How It Works

Token: has token all const of my language

Scanner: converts words into tokens

Parser: gives meaning to token

Analyzer: checks if the source func needs injection

Rewriter: writes the injection

##Need
```
Go 1.22+
```
## Run
```Go
go run .
```
