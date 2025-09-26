package prototype

import (
	"encoding/json"
	"fmt"
)

type ObjectPrototype struct {
	Fields     map[string]string
	Flags      map[string]bool
	Properties map[string]float32
}

func NewObjectPrototypeFromJSON(jsonStr string) (*ObjectPrototype, error) {
	var raw map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &raw); err != nil {
		return nil, err
	}

	obj := &ObjectPrototype{
		Fields:     make(map[string]string),
		Flags:      make(map[string]bool),
		Properties: make(map[string]float32),
	}

	for k, v := range raw {
		switch val := v.(type) {
		case string:
			obj.Fields[k] = val
		case bool:
			obj.Flags[k] = val
		case float64:
			obj.Properties[k] = float32(val)
		default:
			// ignore other types
			fmt.Printf("Unsupported type for key %s\n", k)
		}
	}

	return obj, nil
}
