package components

import (
	"mud-engine/internal/world/entity/component"
	"mud-engine/internal/world/entity/prototype"
)

// Can_Wear enables Wear verb
type CanWear struct {
	component.OComponent
}

func (c *CanWear) BuildComponent(p *prototype.ObjectPrototype) component.OComponent {
	newComponent := CanWear{}
	oComponent := newComponent.OComponent

	if val, ok := p.Flags["can_wear"]; ok {
		oComponent.SetFlag("can_wear", val)
	} else {
		oComponent.SetFlag("can_wear", false)
	}

	if val, ok := p.Fields["wear_location"]; ok {
		oComponent.SetField("wear_location", val)
	} else {
		oComponent.SetField("wear_location", "body")
	}

	if val, ok := p.Properties["armor_class"]; ok {
		oComponent.SetProperty("armor_class", val)
	} else {
		oComponent.SetProperty("armor_class", 0.0)
	}

	if val, ok := p.Fields["material"]; ok {
		oComponent.SetField("material", val)
	} else {
		oComponent.SetField("material", "cloth")
	}

	return newComponent.OComponent
}