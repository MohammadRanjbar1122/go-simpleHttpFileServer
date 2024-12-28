# HttpFileServer

### A simple go software and program to build a http file server 📂 🌐

**Please give it a stars if you like this simple program (:**

![License](https://img.shields.io/badge/license-MIT-blue) ![Contributors](https://img.shields.io/github/contributors/MohammadRanjbar1122/go-simpleHttpFileServer) ![Stars](https://img.shields.io/github/stars/MohammadRanjbar1122/go-simpleHttpFileServer?style=social)

## Features ✨

* Domain setting
* Fast and optimal
* Can be developed and updated
* Ability to set port and path

## Setup 🔧

You can also download the binary and executable version of the program from [releases](https://github.com/MohammadRanjbar1122/go-simpleHttpFileServer/releases/) and you can run it using the following commands with go cli:

```shell
go mod tidy
go run main.go
```

`You can also specify the path and port and domain of the server. You can do this by setting the environment variables in the terminal `:

### on Linux :
```bash
PORT=8000 PATH_DIR=/ DOMAIN_L=custom.mx go run main.go 
```

### on Windows :
```shell
set PATH_DIR=/
set PORT=8000
set DOMAIN_L=custom.mx
```

`If you want to build a binary executable from go codes, enter this command in the terminal :`

### on Linux :
```bash
GOOS=linux GOARCH=386 go build -o 32-64-windowsHttpFileServer.exe -ldflags="-s -w" main.go
```

### on Windows :
```shell
set GOOS=windows
set GOARCH=386
go build -o 32-64-windowsHttpFileServer.exe -ldflags="-s -w" main.go
```
