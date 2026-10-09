# jarvis-sandbox
Sandbox for Jarvis coding sessions (jarvis-code end-to-end tests)

## Usage

`hello` is a small command-line program that prints a greeting.

```sh
go run ./cmd/hello                # Hello, world!
go run ./cmd/hello -name Gopher   # Hello, Gopher!
```

Build or install a `hello` binary:

```sh
go build -o hello ./cmd/hello && ./hello -name Gopher
go install github.com/Deniel9204/jarvis-sandbox/cmd/hello@latest
```

Run the tests:

```sh
go test ./...
```
