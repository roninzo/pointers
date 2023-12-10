# Pointers package

Golang pointer helper.


## Example

```go
import "pointers"

p := pointers.New(25)

if *p == 25 {
    fmt.Printf("p is a pointer to int %d\n", 25)
}
```

## Installation

```bash
go get github.com/roninzo/pointers
```

## Documentation

See code comments.

## Requirement

Go version >= 1.20

## Tests results

```bash
% go test -v -race -timeout 30s -coverprofile=/var/folders/c1/t54lmvkx04715z80v5scjc6c0000gn/T/vscode-goSZW4FX/go-code-cover github.com/roninzo/pointers
=== RUN   TestNew
--- PASS: TestNew (0.00s)
PASS
coverage: 100.0% of statements
ok      github.com/roninzo/pointers     1.250s  coverage: 100.0% of statements
```

## Articles

- [Using generics to get a pointer to any type, in Go](https://www.jvt.me/posts/2022/07/29/go-pointer-generic/)

## Alternatives

- https://github.com/xorcare/pointer
- https://github.com/carlmjohnson/new

## License

[![license](https://img.shields.io/badge/license-MIT-green "The MIT License (MIT)")](LICENSE)
=======
