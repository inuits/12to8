TEST ?= unit

test: build
ifeq (acceptance,${TEST})
	cp 925r.yml 925r
	cd 925r && docker build -t 925r:upstream -f scripts/docker/Dockerfile .
endif
	./tests/${TEST}.sh

build:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure
	go build
