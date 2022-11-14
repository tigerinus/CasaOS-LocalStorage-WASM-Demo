#!/usr/bin/bash

BUILD_DIR=$(dirname "${BASH_SOURCE[0]}")/dist/demo_js_wasm

cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" "$BUILD_DIR"
