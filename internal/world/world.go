package world

import (
	"mud-engine/internal/world/entity"
	"mud-engine/internal/world/entity/prototype"
	"mud-engine/internal/world/entity/verb"
	"mud-engine/internal/world/event"
)

type GameWorld struct {
	objects      map[string]entity.WorldObject
	verbs        []verb.VerbType
	scopes       map[string]interface{}
	eventHandler *event.EventHandler
}

func (gw *GameWorld) add_object(id string, parameters prototype.ObjectPrototype)
	obj := entity.NewWorldObject(id, "", "", "", nil, nil)

	if gw.objects == nil {
		gw.objects = make(map[string]entity.WorldObject)
	}
	gw.objects[id] = *obj


}

func (gw *GameWorld) object_by_id(id string) (entity.WorldObject, bool) {
	obj, ok := gw.objects[id]
	return obj, ok
}

func (gw *GameWorld) register_verb(verb verb.VerbType) {
	gw.verbs = append(gw.verbs, verb)
}
