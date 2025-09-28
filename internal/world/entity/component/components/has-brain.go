package components

import (
	"mud-engine/internal/world/entity/component"
	"mud-engine/internal/world/entity/prototype"
)

// Has_Brain enables Consider verb and with eyes enables Close Examine
type HasBrain struct {
	component.OComponent
}

func (h *HasBrain) BuildComponent(p *prototype.ObjectPrototype) component.OComponent {
	newComponent := HasBrain{}
	oComponent := newComponent.OComponent

	// Initialize maps
	oComponent.Flags = make(map[string]bool)
	oComponent.Fields = make(map[string]string)
	oComponent.Properties = make(map[string]float32)

	if val, ok := p.Flags["has_brain"]; ok {
		oComponent.SetFlag("has_brain", val)
	} else {
		oComponent.SetFlag("has_brain", false)
	}

	if val, ok := p.Properties["intelligence"]; ok {
		oComponent.SetProperty("intelligence", val)
	} else {
		oComponent.SetProperty("intelligence", 10.0)
	}

	if val, ok := p.Fields["brain_type"]; ok {
		oComponent.SetField("brain_type", val)
	} else {
		oComponent.SetField("brain_type", "thinking brain")
	}

	return newComponent.OComponent
}