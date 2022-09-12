GOPATH:=$(shell go env GOPATH)
VERSION=0.0.1

.PHONY: build

build:
	    export GO111MODULE=on
		go build -o ./test-exporter ./main.go

.PHONY: docker

images:
	    docker build -f ./Dockerfile -t test-exporter:${VERSION} .

publish:
	    docker push test-exporter:${VERSION}
