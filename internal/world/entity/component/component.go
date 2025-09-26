// Package component provides the component system for WikiMUD Engine's ECS architecture.
//
// Components are modular pieces of data and behavior that can be attached to entities.
// This package defines the base OComponent type and the ComponentManager that
// orchestrates component lifecycle and access.
package component

import "mud-engine/internal/world/entity/prototype"

// OComponent is the base component type that all components embed.
// It provides the fundamental data storage and manipulation methods
// for the three types of component data: Flags, Fields, and Properties.
//
// - Flags: Boolean values (true/false states)
// - Fields: String values (names, descriptions, identifiers)  
// - Properties: Numeric values (stats, quantities, measurements)
type OComponent struct {
	Name       string                   // Human-readable component name
	Flags      map[string]bool          // Boolean component state
	Fields     map[string]string        // String component data
	Properties map[string]float32       // Numeric component data
	Initialize func(c *OComponent, o *prototype.ObjectPrototype) error // Component initialization
	Refresh    func(c *OComponent) error // Component refresh/update logic
}

// SetProperty sets a numeric property on this component.
// Properties are used for quantitative data like stats, counts, or measurements.
//
// Parameters:
//   - key: The property identifier
//   - value: The numeric value to store
func (c *OComponent) SetProperty(key string, value float32) {
	c.Properties[key] = value
}

// GetProperty retrieves a numeric property from this component.
// Returns the property value and a boolean indicating whether the property exists.
//
// Parameters:
//   - key: The property identifier to retrieve
//
// Returns:
//   - float32: The property value (0 if not found)
//   - bool: True if the property exists, false otherwise
func (c *OComponent) GetProperty(key string) (float32, bool) {
	val, ok := c.Properties[key]
	if !ok {
		return 0, false
	}
	return val, true
}

// SetField sets a string field on this component.
// Fields are used for textual data like names, descriptions, or identifiers.
//
// Parameters:
//   - key: The field identifier
//   - value: The string value to store
func (c *OComponent) SetField(key string, value string) {
	c.Fields[key] = value
}

// GetField retrieves a string field from this component.
// Returns the field value and a boolean indicating whether the field exists.
//
// Parameters:
//   - key: The field identifier to retrieve
//
// Returns:
//   - string: The field value (empty string if not found)
//   - bool: True if the field exists, false otherwise
func (c *OComponent) GetField(key string) (string, bool) {
	val, ok := c.Fields[key]
	if !ok {
		return "", false
	}
	return val, true
}

// SetFlag sets a boolean flag on this component.
// Flags are used for binary states like enabled/disabled or present/absent.
//
// Parameters:
//   - key: The flag identifier
//   - value: The boolean value to store
func (c *OComponent) SetFlag(key string, value bool) {
	c.Flags[key] = value
}

// GetFlag retrieves a boolean flag from this component.
// Returns the flag value and a boolean indicating whether the flag exists.
//
// Parameters:
//   - key: The flag identifier to retrieve
//
// Returns:
//   - bool: The flag value (false if not found)
//   - bool: True if the flag exists, false otherwise
func (c *OComponent) GetFlag(key string) (bool, bool) {
	val, ok := c.Flags[key]
	if !ok {
		return false, false
	}
	return val, true
}

// BuildComponent creates a component instance from prototype data.
// This is the default implementation that should be overridden by specific components.
//
// Specific component types should embed OComponent and override this method
// to implement their own initialization logic based on prototype data.
//
// Parameters:
//   - proto: The object prototype containing initialization data
//
// Returns:
//   - *OComponent: Pointer to the initialized component (nil in base implementation)
func (c *OComponent) BuildComponent(proto *prototype.ObjectPrototype) *OComponent {
	return nil
}
