package components

import (
	"mud-engine/internal/world/entity/component"
	"mud-engine/internal/world/entity/prototype"
)

type CanOpen struct {
	component.OComponent
}

func (c *CanOpen) BuildComponent(p *prototype.ObjectPrototype) component.OComponent {
	newComponent := CanOpen{}
	oComponent := newComponent.OComponent

	if val, ok := p.Flags["can_open"]; ok {
		oComponent.SetFlag("can_open", val)
	} else {
		oComponent.SetFlag("can_open", false)
	}

	if val, ok := p.Flags["can_lock"]; ok {
		oComponent.SetFlag("can_lock", val)
	} else {
		oComponent.SetFlag("can_lock", false)
	}

	if val, ok := p.Flags["can_lock"]; ok && val {
		oComponent.SetFlag("is_locked", p.Flags["is_locked"])
	} else {
		oComponent.SetFlag("is_locked", true)
	}

	if val, ok := p.Fields["key"]; ok {
		oComponent.SetField("key", val)
	} else {
		oComponent.SetField("key", "")
	}

	return newComponent.OComponent
}
