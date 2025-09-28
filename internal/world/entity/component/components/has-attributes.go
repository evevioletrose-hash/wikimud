package components

import (
	"mud-engine/internal/world/entity/component"
	"mud-engine/internal/world/entity/prototype"
)

// Has_Attributes gives an object key value pairs that contain attributes meant to be rolled against
type HasAttributes struct {
	component.OComponent
}

func (h *HasAttributes) BuildComponent(p *prototype.ObjectPrototype) component.OComponent {
	newComponent := HasAttributes{}
	oComponent := newComponent.OComponent

	if val, ok := p.Flags["has_attributes"]; ok {
		oComponent.SetFlag("has_attributes", val)
	} else {
		oComponent.SetFlag("has_attributes", false)
	}

	// Common RPG attributes
	if val, ok := p.Properties["strength"]; ok {
		oComponent.SetProperty("strength", val)
	} else {
		oComponent.SetProperty("strength", 10.0)
	}

	if val, ok := p.Properties["dexterity"]; ok {
		oComponent.SetProperty("dexterity", val)
	} else {
		oComponent.SetProperty("dexterity", 10.0)
	}

	if val, ok := p.Properties["constitution"]; ok {
		oComponent.SetProperty("constitution", val)
	} else {
		oComponent.SetProperty("constitution", 10.0)
	}

	if val, ok := p.Properties["wisdom"]; ok {
		oComponent.SetProperty("wisdom", val)
	} else {
		oComponent.SetProperty("wisdom", 10.0)
	}

	if val, ok := p.Properties["charisma"]; ok {
		oComponent.SetProperty("charisma", val)
	} else {
		oComponent.SetProperty("charisma", 10.0)
	}

	return newComponent.OComponent
}