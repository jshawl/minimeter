.PHONY: .dev .test
dev:
	find . -name '*.go' | entr -r go run ./cmd/minimeter
test:
	go test -v ./...