TEST ?= unit
VERSION ?= locked

test: build
	./tests/${TEST}.sh

build:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure
ifeq (latest,${VERSION})
	dep ensure -update
	git diff
endif
	go build
