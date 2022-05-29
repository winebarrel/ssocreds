.PHONY: all
all: vet build

.PHONY: build
build:
	go build ./cmd/ssocreds

.PHONY: vet
vet:
	go vet ./...

.PHONY: clean
clean:
	rm -f ssocreds ssocreds.exe
