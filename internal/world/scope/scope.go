// Package scope provides entity management within specific contexts.
//
// Scopes are containers that manage entities sharing a common context,
// such as a room, area, or zone in a MUD. They provide entity organization,
// lookup capabilities, and serve as the foundation for rendering and
// event routing systems.
package scope

import (
	"mud-engine/internal/world/entity"
)

// Scope represents a contextual container for entities.
// 
// In MUD terminology, a scope typically represents:
// - A room where players and objects exist
// - An area containing multiple related rooms  
// - A zone with specific themes or properties
// - Any logical grouping of entities that share context
//
// Scopes are renderable and used for trigger routing, making them
// central to the MUD's spatial and logical organization.
type Scope struct {
	entities map[string]entity.WorldObject // Map of entity IDs to entities
}

// NewScope creates a new empty scope ready to contain entities.
// 
// Returns:
//   - *Scope: A new scope with initialized entity storage
func NewScope() *Scope {
	return &Scope{
		entities: make(map[string]entity.WorldObject),
	}
}

// AddEntity adds an entity to this scope under the specified ID.
// 
// If an entity with the same ID already exists, it will be replaced.
// Entity IDs should be unique within the scope to prevent conflicts.
//
// Parameters:
//   - id: Unique identifier for the entity within this scope
//   - obj: The WorldObject entity to add
func (s *Scope) AddEntity(id string, obj entity.WorldObject) {
	s.entities[id] = obj
}

// RemoveEntity removes an entity from this scope by ID.
// 
// If the entity doesn't exist, this operation does nothing.
// This is useful for implementing pickup/drop mechanics, entity
// destruction, or movement between scopes.
//
// Parameters:
//   - id: The identifier of the entity to remove
func (s *Scope) RemoveEntity(id string) {
	delete(s.entities, id)
}

// GetEntity retrieves an entity from this scope by ID.
// 
// Returns the entity and a boolean indicating whether it was found.
// This is the primary method for entity lookup within scopes.
//
// Parameters:
//   - id: The identifier of the entity to retrieve
//
// Returns:
//   - entity.WorldObject: The entity (zero value if not found)
//   - bool: True if the entity exists, false otherwise
func (s *Scope) GetEntity(id string) (entity.WorldObject, bool) {
	obj, ok := s.entities[id]
	return obj, ok
}
