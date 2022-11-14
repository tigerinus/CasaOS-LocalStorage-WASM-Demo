#!/usr/bin/bash

set -e

BUILD_DIR=$(dirname "${BASH_SOURCE[0]}")/build
DIST_DIR=$(dirname "${BASH_SOURCE[0]}")/dist

rm -rf "$DIST_DIR"
mkdir -p "$DIST_DIR"

go generate
go mod tidy
tinygo build -o dist/local_storage.wasm -target wasm main.go 

cp -rv "$(go env GOROOT)/misc/wasm/wasm_exec.js" "$DIST_DIR"
cp -rv "$BUILD_DIR"/* "$DIST_DIR"