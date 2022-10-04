#!/usr/bin/env bash

# https://freshman.tech/snippets/go/cross-compile-go-programs/
# -------------------------------------------------------------


# == LINUX ==

echo ">> linux"
# 32-bit
docker run --rm -v "${PWD}":/usr/src/app -w /usr/src/app -e GOARCH=386 golang:1.19 go build -v -o ./build/logwatch-386-linux ./cmd/logwatch
# 64-bit
docker run --rm -v "${PWD}":/usr/src/app -w /usr/src/app -e GOARCH=amd64 golang:1.19 go build -v -o ./build/logwatch-amd64-linux ./cmd/logwatch
echo "<< linux"



# == WIDNOWS ==

echo ">> windows"
# 32-bit
docker run --rm -v "${PWD}":/usr/src/app -w /usr/src/app -e GOOS=windows -e GOARCH=386 golang:1.19 go build -v -o ./build/logwatch-386.exe ./cmd/logwatch
# 64-bit
docker run --rm -v "${PWD}":/usr/src/app -w /usr/src/app -e GOOS=windows -e GOARCH=amd64 golang:1.19 go build -v -o ./build/logwatch-amd64.exe ./cmd/logwatch
echo "<< windows"



# == MACOS ==

echo ">> macos"
# 32-bit
docker run --rm -v "${PWD}":/usr/src/app -w /usr/src/app -e GOOS=darwin -e GOARCH=386 golang:1.19 go build -v -o ./build/logwatch-386-darwin ./cmd/logwatch
# 64-bit
docker run --rm -v "${PWD}":/usr/src/app -w /usr/src/app -e GOOS=darwin -e GOARCH=amd64 golang:1.19 go build -v -o ./build/logwatch-amd64-darwin ./cmd/logwatch
echo "<< macos"