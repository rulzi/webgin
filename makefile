GO111MODULE=on

build:
	export GO111MODULE on; \
	go build ./...

build-generate:
	export GO111MODULE on; \
	go build -o webgin webgin/cmd