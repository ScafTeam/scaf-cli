.PHONY: all
BUILD_VERSION:=$(shell git describe --tags)

all: scaf.out scaf.exe

# golang for linux
scaf.out: main.go
	echo "build version: $(BUILD_VERSION)"
	env GOOS=linux GOARCH=amd64 go build -o build/scaf-$(BUILD_VERSION)-linux-amd64.out main.go

# golang for windows
scaf.exe: main.go
	env GOOS=windows GOARCH=amd64 go build -o build/scaf-$(BUILD_VERSION)-windows-amd64.exe main.go

# clean
clean:
	rm -rf build/*
