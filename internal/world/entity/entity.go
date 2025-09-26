package entity

import (
	"mud-engine/internal/world/entity/component"
	"mud-engine/internal/world/entity/prototype"
)

type WorldObject struct {
	id          string
	name        string
	description string
	components  *component.ComponentManager
	namespace   string
	parent      *WorldObject
}

func (wo *WorldObject) EmitEvent(event string, params map[string]interface{}) {

}

func (wo *WorldObject) HandleEvent(event string, params map[string]interface{}) {
	// Default implementation: do nothing.
	// You can extend this to handle specific events or propagate further.
}

func NewEntity(id, prototype *prototype.ObjectPrototype) *WorldObject {
	return &WorldObject{
		id:          id,
		name:        name,
		description: description,
		components:  component.NewComponentManager(prototype),
		namespace:   namespace,
		parent:      parent,
	}
}

func (wo *WorldObject) ID() string {
	return wo.id
}

func (wo *WorldObject) Name() string {
	return wo.name
}

func (wo *WorldObject) Description() string {
	return wo.description
}

func (wo *WorldObject) Components() []component.OComponent {
	return wo.components.GetComponents()
}

func (wo *WorldObject) Namespace() string {
	return wo.namespace
}

func (wo *WorldObject) Parent() *WorldObject {
	return wo.parent
}
