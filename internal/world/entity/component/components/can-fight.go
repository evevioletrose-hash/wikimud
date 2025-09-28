package components

import (
	"mud-engine/internal/world/entity/component"
	"mud-engine/internal/world/entity/prototype"
)

// Can_Fight enables combat capabilities
type CanFight struct {
	component.OComponent
}

func (c *CanFight) BuildComponent(p *prototype.ObjectPrototype) component.OComponent {
	newComponent := CanFight{}
	oComponent := newComponent.OComponent

	if val, ok := p.Flags["can_fight"]; ok {
		oComponent.SetFlag("can_fight", val)
	} else {
		oComponent.SetFlag("can_fight", false)
	}

	if val, ok := p.Properties["health"]; ok {
		oComponent.SetProperty("health", val)
	} else {
		oComponent.SetProperty("health", 100.0)
	}

	if val, ok := p.Properties["attack_power"]; ok {
		oComponent.SetProperty("attack_power", val)
	} else {
		oComponent.SetProperty("attack_power", 10.0)
	}

	if val, ok := p.Properties["defense"]; ok {
		oComponent.SetProperty("defense", val)
	} else {
		oComponent.SetProperty("defense", 5.0)
	}

	if val, ok := p.Fields["combat_style"]; ok {
		oComponent.SetField("combat_style", val)
	} else {
		oComponent.SetField("combat_style", "defensive")
	}

	return newComponent.OComponent
}