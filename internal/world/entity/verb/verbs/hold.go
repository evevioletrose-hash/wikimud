package verbs

import (
	"mud-engine/internal/world/entity"
	"mud-engine/internal/world/entity/verb"
)

// Hold attempts to hold an object in hand
// Held items are used with actions. When fighting your held items attributes are used in calculations.
type Hold struct {
	verb.VerbType
}

func (h *Hold) Execute(actor entity.WorldObject, targets []entity.WorldObject, params []string) string {
	if len(targets) == 0 {
		return "Hold what?"
	}
	
	hand := "right"
	if len(params) > 0 {
		hand = params[0]
	}
	
	return "You attempt to hold " + targets[0].Name() + " in your " + hand + " hand."
}

// GetRequiredComponents returns components needed for this verb
func (h *Hold) GetRequiredComponents() []string {
	return []string{"HasHands", "CanHold"}
}