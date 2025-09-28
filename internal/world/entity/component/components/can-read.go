package components

import (
	"mud-engine/internal/world/entity/component"
	"mud-engine/internal/world/entity/prototype"
)

// Can_Read enables Read verb and contains book text in the form of topics and text blocks
type CanRead struct {
	component.OComponent
}

func (c *CanRead) BuildComponent(p *prototype.ObjectPrototype) component.OComponent {
	newComponent := CanRead{}
	oComponent := newComponent.OComponent

	// Initialize maps
	oComponent.Flags = make(map[string]bool)
	oComponent.Fields = make(map[string]string)
	oComponent.Properties = make(map[string]float32)

	if val, ok := p.Flags["can_read"]; ok {
		oComponent.SetFlag("can_read", val)
	} else {
		oComponent.SetFlag("can_read", false)
	}

	if val, ok := p.Fields["content"]; ok {
		oComponent.SetField("content", val)
	} else {
		oComponent.SetField("content", "")
	}

	if val, ok := p.Fields["topics"]; ok {
		oComponent.SetField("topics", val)
	} else {
		oComponent.SetField("topics", "")
	}

	if val, ok := p.Fields["language"]; ok {
		oComponent.SetField("language", val)
	} else {
		oComponent.SetField("language", "common")
	}

	return newComponent.OComponent
}