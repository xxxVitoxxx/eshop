lint:
	golangci-lint run

test:
	go test ./... -count=1

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out
	rm coverage.out

build:
	GOOS=linux&SET GOARCH=arm64 go build -o build/eshop ./cmd/main.go

build-mac:
	GOOS=darwin GOARCH=amd64 go build -o build/eshop ./cmd/main.go

clean:
	rm -r build

.PHONY: lint test coverage build build-mac clean
