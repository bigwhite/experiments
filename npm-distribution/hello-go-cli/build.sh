#!/usr/bin/env bash
set -e
mkdir -p bin
build() {
  GOOS=$1 GOARCH=$2 CGO_ENABLED=0 \
    go build -ldflags="-s -w" -o "bin/hello-$1-$2$3" .
}
build darwin  arm64
build darwin  amd64
build linux   amd64
build linux   arm64
build windows amd64 .exe
