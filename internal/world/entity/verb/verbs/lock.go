package verbs

import (
	"mud-engine/internal/world/entity"
	"mud-engine/internal/world/entity/verb"
)

// Lock prevents an object from being opened without a key
type Lock struct {
	verb.VerbType
}

func (l *Lock) Execute(actor entity.WorldObject, targets []entity.WorldObject, params []string) string {
	if len(targets) == 0 {
		return "Lock what?"
	}
	
	if len(params) == 0 {
		return "Lock " + targets[0].Name() + " with what?"
	}
	
	key := params[0]
	return "You attempt to lock " + targets[0].Name() + " with " + key + "."
}

// GetRequiredComponents returns components needed for this verb
func (l *Lock) GetRequiredComponents() []string {
	return []string{"HasHands"}
}

// Unlock tries to use an object as a key to unlock an object
type Unlock struct {
	verb.VerbType
}

func (u *Unlock) Execute(actor entity.WorldObject, targets []entity.WorldObject, params []string) string {
	if len(targets) == 0 {
		return "Unlock what?"
	}
	
	if len(params) == 0 {
		return "Unlock " + targets[0].Name() + " with what?"
	}
	
	key := params[0]
	return "You attempt to unlock " + targets[0].Name() + " with " + key + "."
}

// GetRequiredComponents returns components needed for this verb
func (u *Unlock) GetRequiredComponents() []string {
	return []string{"HasHands"}
}