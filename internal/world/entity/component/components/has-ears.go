package components

import (
	"mud-engine/internal/world/entity/component"
	"mud-engine/internal/world/entity/prototype"
)

// Has_Ears enables Listen verb and receive heard events in scope
type HasEars struct {
	component.OComponent
}

func (h *HasEars) BuildComponent(p *prototype.ObjectPrototype) component.OComponent {
	newComponent := HasEars{}
	oComponent := newComponent.OComponent

	if val, ok := p.Flags["has_ears"]; ok {
		oComponent.SetFlag("has_ears", val)
	} else {
		oComponent.SetFlag("has_ears", false)
	}

	if val, ok := p.Properties["hearing_acuity"]; ok {
		oComponent.SetProperty("hearing_acuity", val)
	} else {
		oComponent.SetProperty("hearing_acuity", 1.0)
	}

	if val, ok := p.Fields["ear_description"]; ok {
		oComponent.SetField("ear_description", val)
	} else {
		oComponent.SetField("ear_description", "keen ears")
	}

	return newComponent.OComponent
}