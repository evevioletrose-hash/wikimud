package components

import (
	"mud-engine/internal/world/entity/component"
	"mud-engine/internal/world/entity/prototype"
)

// Has_Exits allows an object to have multiple exits. Exits can be hidden based on object conditions
type HasExits struct {
	component.OComponent
}

func (h *HasExits) BuildComponent(p *prototype.ObjectPrototype) component.OComponent {
	newComponent := HasExits{}
	oComponent := newComponent.OComponent

	if val, ok := p.Flags["has_exits"]; ok {
		oComponent.SetFlag("has_exits", val)
	} else {
		oComponent.SetFlag("has_exits", false)
	}

	if val, ok := p.Fields["north"]; ok {
		oComponent.SetField("north", val)
	} else {
		oComponent.SetField("north", "")
	}

	if val, ok := p.Fields["south"]; ok {
		oComponent.SetField("south", val)
	} else {
		oComponent.SetField("south", "")
	}

	if val, ok := p.Fields["east"]; ok {
		oComponent.SetField("east", val)
	} else {
		oComponent.SetField("east", "")
	}

	if val, ok := p.Fields["west"]; ok {
		oComponent.SetField("west", val)
	} else {
		oComponent.SetField("west", "")
	}

	if val, ok := p.Fields["up"]; ok {
		oComponent.SetField("up", val)
	} else {
		oComponent.SetField("up", "")
	}

	if val, ok := p.Fields["down"]; ok {
		oComponent.SetField("down", val)
	} else {
		oComponent.SetField("down", "")
	}

	if val, ok := p.Fields["hidden_exits"]; ok {
		oComponent.SetField("hidden_exits", val)
	} else {
		oComponent.SetField("hidden_exits", "")
	}

	return newComponent.OComponent
}