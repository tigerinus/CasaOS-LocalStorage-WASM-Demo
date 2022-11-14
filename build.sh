#!/usr/bin/bash

set -e

BASE_DIR=$(dirname "${BASH_SOURCE[0]}")
pushd "$BASE_DIR"

go generate
go mod tidy

rm -rf ./dist/*
mkdir -p ./dist

tinygo build -o dist/local_storage.wasm -target wasm main.go 

cp -rv "$(tinygo env TINYGOROOT)/targets/wasm_exec.js" dist/
cp -rv build/* dist/

popd