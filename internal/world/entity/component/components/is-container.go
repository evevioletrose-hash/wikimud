package components

import (
	"mud-engine/internal/world/entity/component"
	"mud-engine/internal/world/entity/prototype"
)

// Is_Container enables placing objects inside of it
type IsContainer struct {
	component.OComponent
}

func (i *IsContainer) BuildComponent(p *prototype.ObjectPrototype) component.OComponent {
	newComponent := IsContainer{}
	oComponent := newComponent.OComponent

	if val, ok := p.Flags["is_container"]; ok {
		oComponent.SetFlag("is_container", val)
	} else {
		oComponent.SetFlag("is_container", false)
	}

	if val, ok := p.Properties["capacity"]; ok {
		oComponent.SetProperty("capacity", val)
	} else {
		oComponent.SetProperty("capacity", 10.0)
	}

	if val, ok := p.Properties["current_load"]; ok {
		oComponent.SetProperty("current_load", val)
	} else {
		oComponent.SetProperty("current_load", 0.0)
	}

	if val, ok := p.Fields["contents"]; ok {
		oComponent.SetField("contents", val)
	} else {
		oComponent.SetField("contents", "")
	}

	return newComponent.OComponent
}