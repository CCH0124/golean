
BINARY_NAME=hello-world

.DEFAULT_GOAL := run
## 產生二進制檔
build: vet
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux main.go
## 產生二進制檔，但在臨時目錄
run:
	go run ./main.go

clean: build
	go clean
	rm ${BINARY_NAME}-linux

fmt: 
	go fmt ./main.go

lint: fmt
	golint ./main.go

vet: fmt
	go vet ./main.go