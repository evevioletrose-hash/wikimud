package verbs

import (
	"mud-engine/internal/world/entity"
	"mud-engine/internal/world/entity/verb"
)

// Look retrieves visual scope based on parameters
type Look struct {
	verb.VerbType
}

func (l *Look) Execute(actor entity.WorldObject, targets []entity.WorldObject, params []string) string {
	// User can define scope by specifying what to look at
	// Multiple objects can be specified
	// Scope is sent to renderer alongside the looking object to render a top level view
	
	if len(targets) == 0 {
		return "You look around the area."
	}
	
	if len(targets) == 1 {
		return "You look at " + targets[0].Name() + "."
	}
	
	result := "You look at "
	for i, target := range targets {
		if i > 0 {
			result += ", "
		}
		result += target.Name()
	}
	return result + "."
}

// GetRequiredComponents returns components needed for this verb
func (l *Look) GetRequiredComponents() []string {
	return []string{"HasEyes"}
}