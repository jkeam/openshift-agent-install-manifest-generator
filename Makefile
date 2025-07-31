.PHONY: build
build: clean test
	GOOS=linux GOARCH=amd64 go build -o oaimg-service main.go

.PHONY: buildmac
buildmac: clean test
	GOOS=darwin GOARCH=arm64 go build -o oaimg-service-arm64 main.go

.PHONY: run
run:
	go run main.go

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	rm -rf ./oaimg-service ./oaimg-service-arm64
