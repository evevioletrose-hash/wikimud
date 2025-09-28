package verbs

import (
	"mud-engine/internal/world/entity"
	"mud-engine/internal/world/entity/verb"
)

// Listen takes the sounds in an active scope and sends them to the renderer
type Listen struct {
	verb.VerbType
}

func (l *Listen) Execute(actor entity.WorldObject, targets []entity.WorldObject, params []string) string {
	if len(targets) == 0 {
		return "You listen carefully to your surroundings."
	}
	
	return "You listen to " + targets[0].Name() + "."
}

// GetRequiredComponents returns components needed for this verb
func (l *Listen) GetRequiredComponents() []string {
	return []string{"HasEars"}
}