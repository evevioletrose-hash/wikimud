// Scopes provide manage entities that share
// a common context, such as a room or area.
// Scopes are renderable and used for trigger
// routing.

package scope

import (
	"mud-engine/internal/world/entity"
)

type Scope struct {
	entities map[string]entity.WorldObject
}

func NewScope() *Scope {
	return &Scope{
		entities: make(map[string]entity.WorldObject),
	}
}

func (s *Scope) AddEntity(id string, obj entity.WorldObject) {
	s.entities[id] = obj
}

func (s *Scope) RemoveEntity(id string) {
	delete(s.entities, id)
}

func (s *Scope) GetEntity(id string) (entity.WorldObject, bool) {
	obj, ok := s.entities[id]
	return obj, ok
}
