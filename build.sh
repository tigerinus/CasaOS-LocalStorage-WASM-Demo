#!/usr/bin/bash

BUILD_DIR=$(dirname "${BASH_SOURCE[0]}")/build

cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" "$BUILD_DIR"/www/