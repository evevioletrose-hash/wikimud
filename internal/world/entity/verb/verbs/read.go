package verbs

import (
	"mud-engine/internal/world/entity"
	"mud-engine/internal/world/entity/verb"
)

// Read attempts to read an object
// Can be given a 'topics' dict that lists keys as headings and values as body text for a book, sign, etc.
type Read struct {
	verb.VerbType
}

func (r *Read) Execute(actor entity.WorldObject, targets []entity.WorldObject, params []string) string {
	if len(targets) == 0 {
		return "Read what?"
	}
	
	return "You attempt to read " + targets[0].Name() + "."
}

// GetRequiredComponents returns components needed for this verb
func (r *Read) GetRequiredComponents() []string {
	return []string{"HasEyes"}
}