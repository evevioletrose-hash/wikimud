package components

import (
	"mud-engine/internal/world/entity/component"
	"mud-engine/internal/world/entity/prototype"
)

// Has_Mouth enables Say, Eat, Drink, Taste verbs
type HasMouth struct {
	component.OComponent
}

func (h *HasMouth) BuildComponent(p *prototype.ObjectPrototype) component.OComponent {
	newComponent := HasMouth{}
	oComponent := newComponent.OComponent

	// Initialize maps
	oComponent.Flags = make(map[string]bool)
	oComponent.Fields = make(map[string]string)
	oComponent.Properties = make(map[string]float32)

	if val, ok := p.Flags["has_mouth"]; ok {
		oComponent.SetFlag("has_mouth", val)
	} else {
		oComponent.SetFlag("has_mouth", false)
	}

	if val, ok := p.Properties["taste_sensitivity"]; ok {
		oComponent.SetProperty("taste_sensitivity", val)
	} else {
		oComponent.SetProperty("taste_sensitivity", 1.0)
	}

	if val, ok := p.Fields["mouth_description"]; ok {
		oComponent.SetField("mouth_description", val)
	} else {
		oComponent.SetField("mouth_description", "expressive mouth")
	}

	return newComponent.OComponent
}