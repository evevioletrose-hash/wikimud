package components

import (
	"mud-engine/internal/world/entity/component"
	"mud-engine/internal/world/entity/prototype"
)

// Can_Drink enables Drink verb
type CanDrink struct {
	component.OComponent
}

func (c *CanDrink) BuildComponent(p *prototype.ObjectPrototype) component.OComponent {
	newComponent := CanDrink{}
	oComponent := newComponent.OComponent

	// Initialize maps
	oComponent.Flags = make(map[string]bool)
	oComponent.Fields = make(map[string]string)
	oComponent.Properties = make(map[string]float32)

	if val, ok := p.Flags["can_drink"]; ok {
		oComponent.SetFlag("can_drink", val)
	} else {
		oComponent.SetFlag("can_drink", false)
	}

	if val, ok := p.Properties["liquid_amount"]; ok {
		oComponent.SetProperty("liquid_amount", val)
	} else {
		oComponent.SetProperty("liquid_amount", 1.0)
	}

	if val, ok := p.Fields["liquid_type"]; ok {
		oComponent.SetField("liquid_type", val)
	} else {
		oComponent.SetField("liquid_type", "water")
	}

	if val, ok := p.Fields["taste"]; ok {
		oComponent.SetField("taste", val)
	} else {
		oComponent.SetField("taste", "refreshing")
	}

	if val, ok := p.Fields["effect"]; ok {
		oComponent.SetField("effect", val)
	} else {
		oComponent.SetField("effect", "")
	}

	return newComponent.OComponent
}