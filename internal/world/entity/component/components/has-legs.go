package components

import (
	"mud-engine/internal/world/entity/component"
	"mud-engine/internal/world/entity/prototype"
)

type HasLegs struct {
	component.OComponent
}

func (h *HasLegs) BuildComponent(p *prototype.ObjectPrototype) component.OComponent {
	newComponent := HasLegs{}
	oComponent := newComponent.OComponent

	if desc, ok := p.Fields["leg_description"]; ok {
		oComponent.SetField("leg_description", desc)
	} else {
		oComponent.SetField("leg_description", "very leglike")
	}

	if count, ok := p.Properties["number_of_legs"]; ok {
		oComponent.SetProperty("leg_count", count)
	} else {
		oComponent.SetProperty("leg_count", 2)
	}
	return newComponent.OComponent
}
