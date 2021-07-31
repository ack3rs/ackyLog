PROJECT := ackyLog
VERSION := $(shell git describe --tag --abbrev=0)
SHA1 := $(shell git rev-parse HEAD)
NOW := $(shell date -u +'%Y%m%d-%H%M%S')


build: 
	go build -o bin/$(PROJECT) -ldflags "-X main.GitCommit=$(SHA1)" -ldflags "-X main.VERSION=$(VERSION)" -ldflags "-X main.BuildTime=$(NOW)"

test:
	@go test  -v -coverprofile=profile.cov ./...
	@go tool cover -func profile.cov

watch: 
	@gow -c run .

run: 
	@go run _examples/basic/basic.go


