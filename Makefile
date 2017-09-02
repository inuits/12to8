TEST ?= unit

test: build
	./tests/${TEST}.sh

build:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure
	go build
