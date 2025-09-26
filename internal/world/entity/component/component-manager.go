package component

import "mud-engine/internal/world/entity/prototype"

type ComponentManager struct {
	Components []OComponent
}

func (cm *ComponentManager) Initialize(prototype *prototype.ObjectPrototype) {

}

func (cm *ComponentManager) GetComponents() []OComponent {
	return cm.Components
}

func NewComponentManager(prototype *prototype.ObjectPrototype) *ComponentManager {
	return &ComponentManager{
		Components: []OComponent{},
	}
}
