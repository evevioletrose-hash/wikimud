package verbs

import (
	"mud-engine/internal/world/entity"
	"mud-engine/internal/world/entity/verb"
	"strings"
)

// Say sends messages into the chat stream for the most local scope in the form of heard events
type Say struct {
	verb.VerbType
}

func (s *Say) Execute(actor entity.WorldObject, targets []entity.WorldObject, params []string) string {
	if len(params) == 0 {
		return "Say what?"
	}
	
	message := strings.Join(params, " ")
	
	if len(targets) == 0 {
		return "You say \"" + message + "\""
	}
	
	return "You say to " + targets[0].Name() + " \"" + message + "\""
}

// GetRequiredComponents returns components needed for this verb
func (s *Say) GetRequiredComponents() []string {
	return []string{"HasMouth"}
}