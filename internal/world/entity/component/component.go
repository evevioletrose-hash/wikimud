package component

import "mud-engine/internal/world/entity/prototype"

type OComponent struct {
	Name       string
	Flags      map[string]bool
	Fields     map[string]string
	Properties map[string]float32
	Initialize func(c *OComponent, o *prototype.ObjectPrototype) error
	Refresh    func(c *OComponent) error
}

func (c *OComponent) SetProperty(key string, value float32) {
	c.Properties[key] = value
}

func (c *OComponent) GetProperty(key string) (float32, bool) {
	val, ok := c.Properties[key]
	if !ok {
		return 0, false
	}
	return val, true
}

func (c *OComponent) SetField(key string, value string) {
	c.Fields[key] = value
}

func (c *OComponent) GetField(key string) (string, bool) {
	val, ok := c.Fields[key]
	if !ok {
		return "", false
	}
	return val, true
}

func (c *OComponent) SetFlag(key string, value bool) {
	c.Flags[key] = value
}

func (c *OComponent) GetFlag(key string) (bool, bool) {
	val, ok := c.Flags[key]
	if !ok {
		return false, false
	}
	return val, true
}

func (c *OComponent) BuildComponent(proto *prototype.ObjectPrototype) *OComponent {
	return nil
}
