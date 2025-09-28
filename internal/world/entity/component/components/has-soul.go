package components

import (
	"mud-engine/internal/world/entity/component"
	"mud-engine/internal/world/entity/prototype"
)

// Has_Soul provides spiritual essence and connection
type HasSoul struct {
	component.OComponent
}

func (h *HasSoul) BuildComponent(p *prototype.ObjectPrototype) component.OComponent {
	newComponent := HasSoul{}
	oComponent := newComponent.OComponent

	// Initialize maps
	oComponent.Flags = make(map[string]bool)
	oComponent.Fields = make(map[string]string)
	oComponent.Properties = make(map[string]float32)

	if val, ok := p.Flags["has_soul"]; ok {
		oComponent.SetFlag("has_soul", val)
	} else {
		oComponent.SetFlag("has_soul", false)
	}

	if val, ok := p.Properties["soul_strength"]; ok {
		oComponent.SetProperty("soul_strength", val)
	} else {
		oComponent.SetProperty("soul_strength", 1.0)
	}

	if val, ok := p.Fields["soul_type"]; ok {
		oComponent.SetField("soul_type", val)
	} else {
		oComponent.SetField("soul_type", "mortal soul")
	}

	return newComponent.OComponent
}