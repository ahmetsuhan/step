ifeq ($(OS),Windows_NT)     # is Windows_NT on XP, 2000, 7, Vista, 10...
	detected_OS := Windows
else
	detected_OS := $(shell uname)  # same as "uname -s"
endif

GOOS=${detected_OS} GOARCH=amd64

help:
	echo "Just make build and run the application"

build:
	go build -o bin/step controller.go struct.go main.go

run:
	go build -o bin/step controller.go struct.go main.go \
	./bin/step