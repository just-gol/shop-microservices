
GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install go-micro.dev/v5/cmd/protoc-gen-micro@v5.12.0
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest

.PHONY: api
api:
	protoc --openapi_out=. --proto_path=. proto/captcha.proto

.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=:. proto/captcha.proto
	
.PHONY: build
build:
	go build -o captcha *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t captcha:latest
