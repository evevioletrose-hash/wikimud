package components

import (
	"mud-engine/internal/world/entity/component"
	"mud-engine/internal/world/entity/prototype"
)

// Has_Physical_Properties defines Height, Width, Composition
type HasPhysicalProperties struct {
	component.OComponent
}

func (h *HasPhysicalProperties) BuildComponent(p *prototype.ObjectPrototype) component.OComponent {
	newComponent := HasPhysicalProperties{}
	oComponent := newComponent.OComponent

	// Initialize maps
	oComponent.Flags = make(map[string]bool)
	oComponent.Fields = make(map[string]string)
	oComponent.Properties = make(map[string]float32)

	if val, ok := p.Flags["has_physical_properties"]; ok {
		oComponent.SetFlag("has_physical_properties", val)
	} else {
		oComponent.SetFlag("has_physical_properties", false)
	}

	if val, ok := p.Properties["height"]; ok {
		oComponent.SetProperty("height", val)
	} else {
		oComponent.SetProperty("height", 1.0)
	}

	if val, ok := p.Properties["width"]; ok {
		oComponent.SetProperty("width", val)
	} else {
		oComponent.SetProperty("width", 1.0)
	}

	if val, ok := p.Properties["weight"]; ok {
		oComponent.SetProperty("weight", val)
	} else {
		oComponent.SetProperty("weight", 1.0)
	}

	if val, ok := p.Fields["composition"]; ok {
		oComponent.SetField("composition", val)
	} else {
		oComponent.SetField("composition", "unknown material")
	}

	return newComponent.OComponent
}