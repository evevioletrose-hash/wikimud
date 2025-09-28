package components

import (
	"mud-engine/internal/world/entity/component"
	"mud-engine/internal/world/entity/prototype"
)

// Can_Eat enables Eat verb
type CanEat struct {
	component.OComponent
}

func (c *CanEat) BuildComponent(p *prototype.ObjectPrototype) component.OComponent {
	newComponent := CanEat{}
	oComponent := newComponent.OComponent

	if val, ok := p.Flags["can_eat"]; ok {
		oComponent.SetFlag("can_eat", val)
	} else {
		oComponent.SetFlag("can_eat", false)
	}

	if val, ok := p.Properties["nutrition_value"]; ok {
		oComponent.SetProperty("nutrition_value", val)
	} else {
		oComponent.SetProperty("nutrition_value", 1.0)
	}

	if val, ok := p.Fields["taste"]; ok {
		oComponent.SetField("taste", val)
	} else {
		oComponent.SetField("taste", "bland")
	}

	if val, ok := p.Fields["effect"]; ok {
		oComponent.SetField("effect", val)
	} else {
		oComponent.SetField("effect", "")
	}

	return newComponent.OComponent
}