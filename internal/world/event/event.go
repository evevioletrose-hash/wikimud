package event

import (
	"mud-engine/internal/world/entity"
	"mud-engine/internal/world/scope"
)

type EventHandler struct {
	channels    map[string]*scope.Scope
	subscribers map[string]map[entity.WorldObject]struct{}
}

func NewEventHandler() *EventHandler {
	return &EventHandler{
		channels:    make(map[string]*scope.Scope),
		subscribers: make(map[string]map[entity.WorldObject]struct{}),
	}
}

// Register a scope as a channel
func (eh *EventHandler) RegisterChannel(name string, s *scope.Scope) {
	eh.channels[name] = s
	if _, ok := eh.subscribers[name]; !ok {
		eh.subscribers[name] = make(map[entity.WorldObject]struct{})
	}
}

// Subscribe a WorldObject to a channel
func (eh *EventHandler) Subscribe(channel string, obj entity.WorldObject) {
	if _, ok := eh.subscribers[channel]; !ok {
		eh.subscribers[channel] = make(map[entity.WorldObject]struct{})
	}
	eh.subscribers[channel][obj] = struct{}{}
}

// Unsubscribe a WorldObject from a channel
func (eh *EventHandler) Unsubscribe(channel string, obj entity.WorldObject) {
	if subs, ok := eh.subscribers[channel]; ok {
		delete(subs, obj)
	}
}

// Send an event to all subscribers of a channel
func (eh *EventHandler) SendEvent(channel string, event interface{}) {
	if subs, ok := eh.subscribers[channel]; ok {
		for obj := range subs {
			obj.HandleEvent(event)
		}
	}
}

// Event represents a generic event in the world.
type Event struct {
	Type    string
	Payload interface{}
}
