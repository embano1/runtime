REPOTAG=embano1/runtime:1.0
BINARY=runtime

default: build

build:
	GOOS=linux CGO_ENABLED=0 go build -a -installsuffix "static" -o ${BINARY}

image:
	docker build -t ${REPOTAG} .

push:
	docker push ${REPOTAG}

clean:
	rm ${BINARY}

all: build image clean push