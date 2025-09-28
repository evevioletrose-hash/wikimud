package verbs

import (
	"mud-engine/internal/world/entity"
	"mud-engine/internal/world/entity/verb"
)

// Wear attempts to wear an object
type Wear struct {
	verb.VerbType
}

func (w *Wear) Execute(actor entity.WorldObject, targets []entity.WorldObject, params []string) string {
	if len(targets) == 0 {
		return "Wear what?"
	}
	
	return "You attempt to wear " + targets[0].Name() + "."
}

// GetRequiredComponents returns components needed for this verb
func (w *Wear) GetRequiredComponents() []string {
	return []string{"HasHands"}
}