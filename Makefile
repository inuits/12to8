TEST ?= unit
VERSION ?= locked

test: build
ifeq (acceptance,${TEST})
ifeq (latest,${VERSION})
	cd 925r && git fetch && git reset --hard origin/master
endif
	cp 925r.yml 925r
	cd 925r && docker build -t 925r:upstream -f scripts/docker/Dockerfile .
endif
	./tests/${TEST}.sh

build:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure
ifeq (latest,${VERSION})
	dep ensure -update
	git diff
endif
	go build
