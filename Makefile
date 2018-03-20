VERSION=1.1
REPO=embano1/runtime
BINARY=runtime

default: build

build:
	GOOS=linux CGO_ENABLED=0 go build -a -installsuffix "static" -o ${BINARY}

image:
	docker build -t ${REPO}:${VERSION} . && docker tag ${REPO}:${VERSION} ${REPO}:latest

push:
	docker push ${REPO} && docker push ${REPO}:latest

clean:
	rm ${BINARY}

all: build image clean push