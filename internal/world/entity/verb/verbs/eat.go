package verbs

import (
	"mud-engine/internal/world/entity"
	"mud-engine/internal/world/entity/verb"
)

// Eat attempts to consume object, render taste and effect
type Eat struct {
	verb.VerbType
}

func (e *Eat) Execute(actor entity.WorldObject, targets []entity.WorldObject, params []string) string {
	if len(targets) == 0 {
		return "Eat what?"
	}
	
	return "You attempt to eat " + targets[0].Name() + "."
}

// GetRequiredComponents returns components needed for this verb
func (e *Eat) GetRequiredComponents() []string {
	return []string{"HasMouth"}
}

// Drink attempts to consume object, render taste and effect
type Drink struct {
	verb.VerbType
}

func (d *Drink) Execute(actor entity.WorldObject, targets []entity.WorldObject, params []string) string {
	if len(targets) == 0 {
		return "Drink what?"
	}
	
	return "You attempt to drink " + targets[0].Name() + "."
}

// GetRequiredComponents returns components needed for this verb
func (d *Drink) GetRequiredComponents() []string {
	return []string{"HasMouth"}
}

// Taste attempts to render object taste
type Taste struct {
	verb.VerbType
}

func (t *Taste) Execute(actor entity.WorldObject, targets []entity.WorldObject, params []string) string {
	if len(targets) == 0 {
		return "Taste what?"
	}
	
	return "You attempt to taste " + targets[0].Name() + "."
}

// GetRequiredComponents returns components needed for this verb
func (t *Taste) GetRequiredComponents() []string {
	return []string{"HasMouth"}
}