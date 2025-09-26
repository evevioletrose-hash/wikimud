package verb

import (
	"fmt"
	"mud-engine/internal/world/entity"
)

type TargetFilter struct {
}

type VerbType struct {
	Name string
}

type VerbHandler struct {
	VerbType *VerbType
}

type Verb struct {
	Name    string
	Actor   entity.WorldObject
	Targets []entity.WorldObject
	Params  []string
	Type    VerbType
}

// Try executes the verb based on the number of targets.
func (v *Verb) Try(actor entity.WorldObject, verbType VerbType, targets []entity.WorldObject, params []string) string {
	v.Actor = actor
	v.Name = verbType.Name
	v.Targets = targets
	v.Params = params

	paramStr := ""
	if len(params) > 0 {
		paramStr = " | Params: " + fmt.Sprint(params)
	}

	switch len(targets) {
	case 0:
		return v.Name + " with no targets." + paramStr
	case 1:
		return v.Name + " with target: " + targets[0].Name() + paramStr
	default:
		names := ""
		for i, t := range targets {
			if i > 0 {
				names += ", "
			}
			names += t.Name()
		}
		return v.Name + " with multiple targets: " + names + paramStr
	}
}

type VerbRules struct {
}
