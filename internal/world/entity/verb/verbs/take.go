package verbs

import (
	"mud-engine/internal/world/entity"
	"mud-engine/internal/world/entity/verb"
)

// Take attempts to hold or add the targeted object to the taking objects container
type Take struct {
	verb.VerbType
}

func (t *Take) Execute(actor entity.WorldObject, targets []entity.WorldObject, params []string) string {
	if len(targets) == 0 {
		return "Take what?"
	}
	
	if len(targets) == 1 {
		return "You attempt to take " + targets[0].Name() + "."
	}
	
	result := "You attempt to take "
	for i, target := range targets {
		if i > 0 {
			result += ", "
		}
		result += target.Name()
	}
	return result + "."
}

// GetRequiredComponents returns components needed for this verb
func (t *Take) GetRequiredComponents() []string {
	return []string{"HasHands"}
}