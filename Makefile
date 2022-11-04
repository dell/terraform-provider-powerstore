
lint:
	echo "Running staticcheck"
	go install honnef.co/go/tools/cmd/staticcheck@latest
	staticcheck ./...
	golint ./...

vet:
	echo "running go vet"
	go vet

fmt:
	gofmt -w -s .

code-check: lint vet fmt

generate:
	go generate

download:
	go mod download

build: download
	mkdir -p out
	go build -v -o ./out

all: download code-check
