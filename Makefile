.PHONY: all
all: vet build

.PHONY: build
build:
	go build ./cmd/ssocreds

.PHONY: vet
vet:
	go vet ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: clean
clean:
	rm -f ssocreds ssocreds.exe
