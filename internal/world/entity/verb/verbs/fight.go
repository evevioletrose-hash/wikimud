package verbs

import (
	"mud-engine/internal/world/entity"
	"mud-engine/internal/world/entity/verb"
)

// Fight attempts to start a fight encounter with object
type Fight struct {
	verb.VerbType
}

func (f *Fight) Execute(actor entity.WorldObject, targets []entity.WorldObject, params []string) string {
	if len(targets) == 0 {
		return "Fight whom?"
	}
	
	return "You attempt to engage " + targets[0].Name() + " in combat."
}

// GetRequiredComponents returns components needed for this verb
func (f *Fight) GetRequiredComponents() []string {
	return []string{"CanFight"}
}

// Consider analyzes an object or situation
type Consider struct {
	verb.VerbType
}

func (c *Consider) Execute(actor entity.WorldObject, targets []entity.WorldObject, params []string) string {
	if len(targets) == 0 {
		return "Consider what?"
	}
	
	return "You consider " + targets[0].Name() + " carefully."
}

// GetRequiredComponents returns components needed for this verb
func (c *Consider) GetRequiredComponents() []string {
	return []string{"HasBrain"}
}

// CloseExamine performs detailed examination
type CloseExamine struct {
	verb.VerbType
}

func (c *CloseExamine) Execute(actor entity.WorldObject, targets []entity.WorldObject, params []string) string {
	if len(targets) == 0 {
		return "Examine what closely?"
	}
	
	return "You examine " + targets[0].Name() + " closely."
}

// GetRequiredComponents returns components needed for this verb
func (c *CloseExamine) GetRequiredComponents() []string {
	return []string{"HasEyes", "HasBrain"}
}