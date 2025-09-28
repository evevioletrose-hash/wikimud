package components

import (
	"mud-engine/internal/world/entity/component"
	"mud-engine/internal/world/entity/prototype"
)

type HasEyes struct {
	component.OComponent
}

func (h *HasEyes) BuildComponent(p *prototype.ObjectPrototype) component.OComponent {
	newComponent := HasEyes{}
	oComponent := newComponent.OComponent

	// Initialize maps
	oComponent.Flags = make(map[string]bool)
	oComponent.Fields = make(map[string]string)
	oComponent.Properties = make(map[string]float32)

	if val, ok := p.Flags["has_eyes"]; ok {
		oComponent.SetFlag("has_eyes", val)
	} else {
		oComponent.SetFlag("has_eyes", false)
	}

	if val, ok := p.Fields["eye_color"]; ok {
		oComponent.SetField("eye_color", val)
	} else {
		oComponent.SetField("eye_color", "")
	}

	if val, ok := p.Properties["eye_count"]; ok {
		oComponent.SetProperty("eye_count", val)
	} else {
		oComponent.SetProperty("eye_count", 2)
	}

	return newComponent.OComponent
}
