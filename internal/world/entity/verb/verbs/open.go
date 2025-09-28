package verbs

import (
	"mud-engine/internal/world/entity"
	"mud-engine/internal/world/entity/verb"
)

// Open attempts to open an object if it can be opened
type Open struct {
	verb.VerbType
}

func (o *Open) Execute(actor entity.WorldObject, targets []entity.WorldObject, params []string) string {
	if len(targets) == 0 {
		return "Open what?"
	}
	
	return "You attempt to open " + targets[0].Name() + "."
}

// GetRequiredComponents returns components needed for this verb
func (o *Open) GetRequiredComponents() []string {
	return []string{"HasHands"}
}

// Close attempts to close an object if it is open
type Close struct {
	verb.VerbType
}

func (c *Close) Execute(actor entity.WorldObject, targets []entity.WorldObject, params []string) string {
	if len(targets) == 0 {
		return "Close what?"
	}
	
	return "You attempt to close " + targets[0].Name() + "."
}

// GetRequiredComponents returns components needed for this verb
func (c *Close) GetRequiredComponents() []string {
	return []string{"HasHands"}
}