#!/bin/bash

# WINDOWS
# 64-bit
GOOS=windows GOARCH=amd64 go build -o build/kubo-socks-0.1.1-windows-amd64.exe -ldflags "-s -w"
# 32-bit
GOOS=windows GOARCH=386 go build -o build/kubo-socks-0.1.1-windows-i386.exe -ldflags "-s -w"

# LINUX
# 64-bit
GOOS=linux GOARCH=amd64 go build -o build/kubo-socks-0.1.1-linux-amd64 -ldflags "-s -w"
# 32-bit
GOOS=linux GOARCH=386 go build -o build/kubo-socks-0.1.1-linux-i386 -ldflags "-s -w"

# DARWIN - MACOS
# 64-bit
GOOS=darwin GOARCH=amd64 go build -o build/kubo-socks-0.1.1-darwin-amd64 -ldflags "-s -w"
# 32-bit
# GOOS=darwin GOARCH=386 go build -o build/kubo-socks-0.1.1-darwin-i386 -ldflags "-s -w"

