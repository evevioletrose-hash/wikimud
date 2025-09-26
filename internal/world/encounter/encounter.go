package encounter

import (
	"mud-engine/internal/world/entity"
	"mud-engine/internal/world/entity/attribute"
	"mud-engine/internal/world/entity/verb"
)

type Encounter struct {
	Participants []*entity.WorldObject
	Location     *entity.WorldObject
	Beats        []EncounterBeat
	WinCons      []*EncounterConditions
	LoseCons     []*EncounterConditions
}

type EncounterConditions struct {
}

type EncounterType struct {
	Name        string
	Description string
}

type EncounterBeat struct {
	Verb        *verb.Verb
	RollWith    *attribute.Attribute
	RollAgainst *attribute.Attribute
}
