GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=codeforces
BINARY_UNIX=$(BINARY_NAME)_unix
BUILD_ENV = env
CODECOV = /bin/bash codecov.sh

all: deps test build

build: build-linux64 build-linux32 build-osx64 build-osx32 build-windows64 build-windows32

test:
	$(CODECOV)
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
deps:
	$(GOGET) github.com/opesun/goquery
	$(GOGET) github.com/bitly/go-simplejson

build-linux64:
	mkdir -p build/linux_64
	${BUILD_ENV} GOARCH=amd64 GOOS=linux ${GOBUILD} -o build/linux_64/${BINARY_NAME} -v

build-osx64:
	mkdir -p build/osx_64
	${BUILD_ENV} GOARCH=amd64 GOOS=darwin ${GOBUILD} -o build/osx_64/${BINARY_NAME} -v

build-windows64:
	mkdir -p build/windows_64
	${BUILD_ENV} GOARCH=amd64 GOOS=windows ${GOBUILD} -o build/windows_64/${BINARY_NAME}.exe -v

build-linux32:
	mkdir -p build/linux_32
	${BUILD_ENV} GOARCH=386 GOOS=linux ${GOBUILD} -o build/linux_32/${BINARY_NAME} -v

build-osx32:
	mkdir -p build/osx_32
	${BUILD_ENV} GOARCH=386 GOOS=darwin ${GOBUILD} -o build/osx_32/${BINARY_NAME} -v

build-windows32:
	mkdir -p build/windows_32
	${BUILD_ENV} GOARCH=386 GOOS=windows ${GOBUILD} -o build/windows_32/${BINARY_NAME}.exe -v
