package components

import (
	"mud-engine/internal/world/entity/component"
	"mud-engine/internal/world/entity/prototype"
)

// Has_Hands enables Take, Hold, Move, Open, Close, Drop, Pickpocket verbs
type HasHands struct {
	component.OComponent
}

func (h *HasHands) BuildComponent(p *prototype.ObjectPrototype) component.OComponent {
	newComponent := HasHands{}
	oComponent := newComponent.OComponent

	if val, ok := p.Flags["has_hands"]; ok {
		oComponent.SetFlag("has_hands", val)
	} else {
		oComponent.SetFlag("has_hands", false)
	}

	if val, ok := p.Properties["hand_count"]; ok {
		oComponent.SetProperty("hand_count", val)
	} else {
		oComponent.SetProperty("hand_count", 2)
	}

	if val, ok := p.Fields["hand_description"]; ok {
		oComponent.SetField("hand_description", val)
	} else {
		oComponent.SetField("hand_description", "dexterous hands")
	}

	return newComponent.OComponent
}