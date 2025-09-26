// Package entity provides the core entity system for WikiMUD Engine.
// 
// This package implements the Entity-Component-System (ECS) architecture,
// where WorldObject represents entities that can have multiple components
// attached to define their behavior and properties.
package entity

import (
	"mud-engine/internal/world/entity/component"
	"mud-engine/internal/world/entity/prototype"
)

// WorldObject represents any entity in the MUD world (players, items, NPCs, rooms, etc.).
// It serves as the foundation of the Entity-Component-System architecture.
//
// WorldObjects are created from prototypes and can have multiple components
// that define their behavior and properties. They exist within scopes and
// can participate in the event system.
type WorldObject struct {
	id          string                        // Unique identifier for this entity
	name        string                        // Display name of the entity
	description string                        // Detailed description text
	components  *component.ComponentManager   // Manages attached components
	namespace   string                        // Logical grouping namespace
	parent      *WorldObject                  // Parent entity for hierarchy
}

// EmitEvent sends an event from this entity to the event system.
// Events are used for loose coupling between game systems.
//
// Parameters:
//   - event: String identifier for the event type
//   - params: Key-value pairs of event-specific data
func (wo *WorldObject) EmitEvent(event string, params map[string]interface{}) {
	// TODO: Implement event emission to global event handler
}

// HandleEvent processes an incoming event targeted at this entity.
// Override this method in specific entity types to implement custom behavior.
//
// The default implementation does nothing, allowing entities to ignore
// events they don't care about.
//
// Parameters:
//   - event: String identifier for the event type
//   - params: Key-value pairs of event-specific data
func (wo *WorldObject) HandleEvent(event string, params map[string]interface{}) {
	// Default implementation: do nothing.
	// You can extend this to handle specific events or propagate further.
}

// NewEntity creates a new WorldObject from a prototype definition.
// This is the primary constructor for entities in the MUD world.
//
// The prototype defines the initial component configuration and properties.
// Components are automatically created and initialized based on the prototype data.
//
// Parameters:
//   - id: Unique identifier for the new entity
//   - prototype: Object prototype defining initial properties and components
//
// Returns:
//   - Pointer to the newly created WorldObject
func NewEntity(id string, prototype *prototype.ObjectPrototype) *WorldObject {
	// Extract name and description from prototype if available
	name := ""
	if protoName, ok := prototype.Fields["name"]; ok {
		name = protoName
	}
	
	description := ""
	if protoDesc, ok := prototype.Fields["description"]; ok {
		description = protoDesc
	}
	
	return &WorldObject{
		id:          id,
		name:        name,
		description: description,
		components:  component.NewComponentManager(prototype),
		namespace:   "", // Default empty namespace
		parent:      nil, // No parent by default
	}
}

// ID returns the unique identifier of this entity.
// IDs should be unique within their scope to prevent conflicts.
func (wo *WorldObject) ID() string {
	return wo.id
}

// Name returns the display name of this entity.
// This is what players will see when the entity is referenced in game text.
func (wo *WorldObject) Name() string {
	return wo.name
}

// Description returns the detailed description of this entity.
// This is typically shown when players examine or look at the entity.
func (wo *WorldObject) Description() string {
	return wo.description
}

// Components returns all components attached to this entity.
// Components define the entity's capabilities and behavior.
func (wo *WorldObject) Components() []component.OComponent {
	return wo.components.GetComponents()
}

// Namespace returns the logical grouping namespace for this entity.
// Namespaces can be used to organize entities by type, area, or function.
func (wo *WorldObject) Namespace() string {
	return wo.namespace
}

// Parent returns the parent entity in the containment hierarchy.
// For example, an item's parent might be a bag or a player's inventory.
// Returns nil if this entity has no parent.
func (wo *WorldObject) Parent() *WorldObject {
	return wo.parent
}
