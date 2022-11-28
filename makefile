.PHONY: all

all: scaf.out scaf.exe

# golang for linux
scaf.out: main.go
	env GOOS=linux GOARCH=amd64 go build -o build/scaf.out main.go

# golang for windows
scaf.exe: main.go
	env GOOS=windows GOARCH=amd64 go build -o build/scaf.exe main.go
