.PHONY: bin
bin:
	mkdir -p bin/

.PHONY: build
build: bin
	go build -o bin/ikea ./cmd/ikea/*.go

.PHONY: mod
mod:
	go mod download

.PHONY: vendor
vendor:
	rm -rf vendor
	go mod vendor

.PHONY: test
test:
	go test ./... -v