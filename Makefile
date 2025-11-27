.PHONY: .dev .test .run
dev:
	find . -name '*.go' | entr -r go run ./cmd/minimeter
test:
	go test -v ./...
run:
	docker build -t minimeter .
	docker run --rm -v minimeter-data:/app/data -p 8080:8080 minimeter
