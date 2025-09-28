package verbs

import (
	"mud-engine/internal/world/entity"
	"mud-engine/internal/world/entity/verb"
)

// Move moves an item from one place to another
type Move struct {
	verb.VerbType
}

func (m *Move) Execute(actor entity.WorldObject, targets []entity.WorldObject, params []string) string {
	if len(targets) == 0 {
		return "Move what?"
	}
	
	if len(params) == 0 {
		return "Move " + targets[0].Name() + " where?"
	}
	
	placement := params[0]
	return "You attempt to move " + targets[0].Name() + " to " + placement + "."
}

// GetRequiredComponents returns components needed for this verb
func (m *Move) GetRequiredComponents() []string {
	return []string{"HasHands"}
}

// Drop places an item on the ground (move overlap)
type Drop struct {
	verb.VerbType
}

func (d *Drop) Execute(actor entity.WorldObject, targets []entity.WorldObject, params []string) string {
	if len(targets) == 0 {
		return "Drop what?"
	}
	
	return "You drop " + targets[0].Name() + "."
}

// GetRequiredComponents returns components needed for this verb
func (d *Drop) GetRequiredComponents() []string {
	return []string{"HasHands"}
}