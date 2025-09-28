package components

import (
	"mud-engine/internal/world/entity/component"
	"mud-engine/internal/world/entity/prototype"
)

// Can_Hold enables Hold verb
type CanHold struct {
	component.OComponent
}

func (c *CanHold) BuildComponent(p *prototype.ObjectPrototype) component.OComponent {
	newComponent := CanHold{}
	oComponent := newComponent.OComponent

	// Initialize maps
	oComponent.Flags = make(map[string]bool)
	oComponent.Fields = make(map[string]string)
	oComponent.Properties = make(map[string]float32)

	if val, ok := p.Flags["can_hold"]; ok {
		oComponent.SetFlag("can_hold", val)
	} else {
		oComponent.SetFlag("can_hold", false)
	}

	if val, ok := p.Fields["held_in_left"]; ok {
		oComponent.SetField("held_in_left", val)
	} else {
		oComponent.SetField("held_in_left", "")
	}

	if val, ok := p.Fields["held_in_right"]; ok {
		oComponent.SetField("held_in_right", val)
	} else {
		oComponent.SetField("held_in_right", "")
	}

	if val, ok := p.Properties["grip_strength"]; ok {
		oComponent.SetProperty("grip_strength", val)
	} else {
		oComponent.SetProperty("grip_strength", 1.0)
	}

	return newComponent.OComponent
}