//go:generate bash -c "mkdir -p codegen/message_bus && go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.12.2 -package message_bus https://raw.githubusercontent.com/IceWhaleTech/CasaOS-MessageBus/main/api/message_bus/openapi.yaml > codegen/message_bus/api.go"

package main

import (
	"wasm-demo/codegen/message_bus"
)

func GetUIProperties(event message_bus.Event) interface{} {
	switch event.SourceID {
	default:
		return nil
	}
}

func main() {
	println("CasaOS-LocalStorage-WASM-Demo loaded.")
}
