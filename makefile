run:
	@go build -o ./bin/todo ./cmd/main.go && ./bin/todo

test:
	@go test ./...