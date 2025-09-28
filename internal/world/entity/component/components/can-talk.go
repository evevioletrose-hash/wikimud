package components

import (
	"mud-engine/internal/world/entity/component"
	"mud-engine/internal/world/entity/prototype"
)

// Can_Talk allows an object to respond when things are said to it
type CanTalk struct {
	component.OComponent
}

func (c *CanTalk) BuildComponent(p *prototype.ObjectPrototype) component.OComponent {
	newComponent := CanTalk{}
	oComponent := newComponent.OComponent

	if val, ok := p.Flags["can_talk"]; ok {
		oComponent.SetFlag("can_talk", val)
	} else {
		oComponent.SetFlag("can_talk", false)
	}

	if val, ok := p.Fields["dialogue_tree"]; ok {
		oComponent.SetField("dialogue_tree", val)
	} else {
		oComponent.SetField("dialogue_tree", "")
	}

	if val, ok := p.Fields["greeting"]; ok {
		oComponent.SetField("greeting", val)
	} else {
		oComponent.SetField("greeting", "Hello there!")
	}

	if val, ok := p.Properties["friendliness"]; ok {
		oComponent.SetProperty("friendliness", val)
	} else {
		oComponent.SetProperty("friendliness", 0.5)
	}

	return newComponent.OComponent
}