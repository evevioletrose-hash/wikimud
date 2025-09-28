package verbs

import (
	"mud-engine/internal/world/entity"
	"mud-engine/internal/world/entity/verb"
)

// Go attempts to move the character in the direction specified
type Go struct {
	verb.VerbType
}

func (g *Go) Execute(actor entity.WorldObject, targets []entity.WorldObject, params []string) string {
	// The going object must have the necessary requirements to travel in the given direction
	// (legs, wings, attributes)
	
	if len(params) == 0 {
		return "Go where?"
	}
	
	direction := params[0]
	return "You attempt to go " + direction + "."
}

// GetRequiredComponents returns components needed for this verb
func (g *Go) GetRequiredComponents() []string {
	return []string{"HasLegs"}
}