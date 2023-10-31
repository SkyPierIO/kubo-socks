#!/bin/bash

version="0.1.3"

# WINDOWS
# 64-bit
GOOS=windows GOARCH=amd64 go build -o build/kubo-socks.exe -ldflags "-s -w" -trimpath
zip build/kubo-socks-${version}-windows-amd64.zip build/kubo-socks.exe LICENSE
rm build/kubo-socks.exe
# 32-bit
GOOS=windows GOARCH=386 go build -o build/kubo-socks.exe -ldflags "-s -w" -trimpath
zip build/kubo-socks-${version}-windows-i386.zip build/kubo-socks.exe LICENSE
rm build/kubo-socks.exe

# LINUX
# 64-bit
GOOS=linux GOARCH=amd64 go build -o build/kubo-socks -ldflags "-s -w" -trimpath
zip build/kubo-socks-${version}-linux-amd64.zip build/kubo-socks LICENSE
rm build/kubo-socks
# 32-bit
GOOS=linux GOARCH=386 go build -o build/kubo-socks -ldflags "-s -w" -trimpath
zip build/kubo-socks-${version}-linux-i386.zip build/kubo-socks LICENSE
rm build/kubo-socks

# DARWIN - MACOS
# 64-bit
GOOS=darwin GOARCH=amd64 go build -o build/kubo-socks -ldflags "-s -w" -trimpath
zip build/kubo-socks-${version}-darwin-amd64.zip build/kubo-socks LICENSE
rm build/kubo-socks
# 32-bit
# GOOS=darwin GOARCH=386 go build -o build/kubo-socks-${version}-darwin-i386 -ldflags "-s -w" -trimpath
# ARM 64-bit (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o build/kubo-socks -ldflags "-s -w" -trimpath
zip build/kubo-socks-${version}-darwin-arm64.zip build/kubo-socks LICENSE
rm build/kubo-socks
