//go:generate bash -c "mkdir -p codegen/message_bus && go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.12.2 -generate types -package message_bus https://raw.githubusercontent.com/IceWhaleTech/CasaOS-MessageBus/main/api/message_bus/openapi.yaml > codegen/message_bus/api.go && go run github.com/mailru/easyjson/...@latest -all model/model.go && go run github.com/mailru/easyjson/...@latest -all codegen/message_bus/api.go"

package main

import (
	"fmt"
	"syscall/js"

	"github.com/mailru/easyjson"

	"wasm-demo/codegen/message_bus"
	"wasm-demo/model"
)

var log = make(chan string)

func PropertiesToMap(properties []message_bus.Property) map[string]string {
	m := make(map[string]string)
	for _, property := range properties {
		m[property.Name] = property.Value
	}

	return m
}

func MapToProperties(m map[string]string) []message_bus.Property {
	properties := make([]message_bus.Property, 0)
	for name, value := range m {
		properties = append(properties, message_bus.Property{
			Name:  name,
			Value: value,
		})
	}

	return properties
}

func GetUIProperties(event message_bus.Event) *model.UIElement {
	var properties map[string]string
	if event.Properties == nil {
		log <- "event.Properties is nil"
		properties = make(map[string]string)
	} else {
		properties = PropertiesToMap(event.Properties)
	}

	switch event.SourceID {

	case "local-storage":
		switch event.Name {

		case "local-storage:disk:added":
			diskVendor, ok := properties["local-storage:vendor"]
			if !ok {
				log <- "local-storage:vendor not found"
				diskVendor = "unknown"
			}

			diskModel, ok := properties["local-storage:model"]
			if !ok {
				log <- "local-storage:model not found"
				diskModel = "unknown"
			}

			return &model.UIElement{
				Title:   "New disk found",
				Icon:    "casaos-icon-disk",
				Message: fmt.Sprintf("A new disk, %s %s, is added", diskVendor, diskModel),
				UIType:  "casaos-ui-notification-style-2",
			}

		case "local-storage:disk:removed":
			diskVendor, ok := properties["local-storage:vendor"]
			if !ok {
				log <- "local-storage:vendor not found"
				diskVendor = "unknown"
			}

			diskModel, ok := properties["local-storage:model"]
			if !ok {
				log <- "local-storage:model not found"
				diskModel = "unknown"
			}

			return &model.UIElement{
				Title:   "Disk removed",
				Icon:    "casaos-icon-disk",
				Message: fmt.Sprintf("A disk, %s %s, is removed", diskVendor, diskModel),
				UIType:  "casaos-ui-notification-style-2",
			}

		default:
			log <- fmt.Sprintf("unknown event name: %s", event.Name)
			return nil
		}

	default:
		log <- fmt.Sprintf("unknown source id: %s", event.SourceID)
		return nil
	}
}

func jsWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			log <- "there should be only 1 argument"
			return nil
		}
		event := message_bus.Event{}
		if err := easyjson.Unmarshal([]byte(args[0].String()), &event); err != nil {
			log <- fmt.Sprintf("failed to unmarshal from argument `%s`: %v", args[0], err)
			return nil
		}
		result := GetUIProperties(event)
		if result == nil {
			log <- "GetUIProperties returned nil"
			return ""
		}
		b, err := easyjson.Marshal(result)
		if err != nil {
			log <- fmt.Sprintf("failed to marshal result: %v", err)
			return ""
		}
		return string(b)
	})
}

func main() {
	defer close(log)

	println("CasaOS-LocalStorage-WASM-Demo loaded.")
	js.Global().Set("GetUIProperties", jsWrapper())

	for msg := range log {
		println(msg)
	}
}
