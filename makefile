.PHONY: all
BUILD_VERSION:=$(shell git describe --tags)
ifneq ($(shell git describe > /dev/null 2>&1 ; echo $$?),0)
BUILD_VERSION:=test
endif

all: scaf.out scaf.exe
	@echo "build version: $(BUILD_VERSION)"

# golang for linux
scaf.out: main.go
	env GOOS=linux GOARCH=amd64 go build -o build/scaf-$(BUILD_VERSION)-linux-amd64.out main.go

# golang for windows
scaf.exe: main.go
	env GOOS=windows GOARCH=amd64 go build -o build/scaf-$(BUILD_VERSION)-windows-amd64.exe main.go

# clean
clean:
	rm -rf build/*
