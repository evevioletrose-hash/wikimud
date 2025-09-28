package verbs

import (
	"mud-engine/internal/world/entity"
	"mud-engine/internal/world/entity/verb"
)

// Pickpocket attempts to covertly take from an object with a container
type Pickpocket struct {
	verb.VerbType
}

func (p *Pickpocket) Execute(actor entity.WorldObject, targets []entity.WorldObject, params []string) string {
	if len(targets) == 0 {
		return "Pickpocket whom?"
	}
	
	return "You attempt to stealthily take something from " + targets[0].Name() + "."
}

// GetRequiredComponents returns components needed for this verb
func (p *Pickpocket) GetRequiredComponents() []string {
	return []string{"HasHands"}
}