# Custom Component Example

This example demonstrates how to create custom components for WikiMUD Engine, extending the entity system with new capabilities.

## Overview

We'll create:
- A **Health** component for combat and healing
- A **Magic** component for spellcasting abilities
- A **Container** component for items that hold other items
- Entities that use these custom components
- Demonstration of component interactions

## Custom Components Implementation

### 1. Health Component

First, let's create a health component for combat systems:

```go
// File: internal/world/entity/component/components/health.go
package components

import (
    "mud-engine/internal/world/entity/component"
    "mud-engine/internal/world/entity/prototype"
)

// Health provides hit points and vitality management for entities.
//
// Properties:
//   - current_health (property): Current hit points
//   - max_health (property): Maximum hit points  
//   - regeneration_rate (property): HP regenerated per time unit
//   - is_alive (flag): Whether the entity is alive
//   - death_message (field): Message displayed when entity dies
type Health struct {
    component.OComponent
}

// BuildComponent initializes a Health component from prototype data.
func (h *Health) BuildComponent(p *prototype.ObjectPrototype) component.OComponent {
    newComponent := Health{}
    oComponent := newComponent.OComponent
    
    // Initialize component name
    oComponent.Name = "Health"
    
    // Initialize maps if they don't exist
    if oComponent.Properties == nil {
        oComponent.Properties = make(map[string]float32)
    }
    if oComponent.Flags == nil {
        oComponent.Flags = make(map[string]bool)
    }
    if oComponent.Fields == nil {
        oComponent.Fields = make(map[string]string)
    }
    
    // Set maximum health
    if maxHP, ok := p.Properties["max_health"]; ok {
        oComponent.SetProperty("max_health", maxHP)
    } else {
        oComponent.SetProperty("max_health", 100.0)
    }
    
    // Set current health (defaults to max if not specified)
    if currentHP, ok := p.Properties["current_health"]; ok {
        oComponent.SetProperty("current_health", currentHP)
    } else {
        maxHP, _ := oComponent.GetProperty("max_health")
        oComponent.SetProperty("current_health", maxHP)
    }
    
    // Set regeneration rate
    if regenRate, ok := p.Properties["regeneration_rate"]; ok {
        oComponent.SetProperty("regeneration_rate", regenRate)
    } else {
        oComponent.SetProperty("regeneration_rate", 1.0)
    }
    
    // Set alive status
    if isAlive, ok := p.Flags["is_alive"]; ok {
        oComponent.SetFlag("is_alive", isAlive)
    } else {
        oComponent.SetFlag("is_alive", true)
    }
    
    // Set death message
    if deathMsg, ok := p.Fields["death_message"]; ok {
        oComponent.SetField("death_message", deathMsg)
    } else {
        oComponent.SetField("death_message", "has fallen")
    }
    
    return newComponent.OComponent
}

// Heal increases the entity's current health by the specified amount.
func (h *Health) Heal(amount float32) {
    current, _ := h.GetProperty("current_health")
    max, _ := h.GetProperty("max_health")
    
    newHealth := current + amount
    if newHealth > max {
        newHealth = max
    }
    
    h.SetProperty("current_health", newHealth)
    
    // Revive if fully healed
    if newHealth > 0 {
        h.SetFlag("is_alive", true)
    }
}

// Damage reduces the entity's current health by the specified amount.
func (h *Health) Damage(amount float32) {
    current, _ := h.GetProperty("current_health")
    newHealth := current - amount
    
    if newHealth <= 0 {
        newHealth = 0
        h.SetFlag("is_alive", false)
    }
    
    h.SetProperty("current_health", newHealth)
}

// IsAlive returns whether the entity is currently alive.
func (h *Health) IsAlive() bool {
    alive, _ := h.GetFlag("is_alive")
    return alive
}
```

### 2. Magic Component

Next, a magic component for spellcasting:

```go
// File: internal/world/entity/component/components/magic.go
package components

import (
    "mud-engine/internal/world/entity/component"
    "mud-engine/internal/world/entity/prototype"
)

// Magic provides magical abilities and mana management for entities.
//
// Properties:
//   - current_mana (property): Current magical energy
//   - max_mana (property): Maximum magical energy
//   - spell_power (property): Effectiveness of spells cast
//   - mana_regeneration (property): Mana regenerated per time unit
//   - can_cast (flag): Whether entity can cast spells
//   - magical_school (field): Primary school of magic
type Magic struct {
    component.OComponent
}

// BuildComponent initializes a Magic component from prototype data.
func (m *Magic) BuildComponent(p *prototype.ObjectPrototype) component.OComponent {
    newComponent := Magic{}
    oComponent := newComponent.OComponent
    
    oComponent.Name = "Magic"
    
    // Initialize maps
    if oComponent.Properties == nil {
        oComponent.Properties = make(map[string]float32)
    }
    if oComponent.Flags == nil {
        oComponent.Flags = make(map[string]bool)
    }
    if oComponent.Fields == nil {
        oComponent.Fields = make(map[string]string)
    }
    
    // Set maximum mana
    if maxMana, ok := p.Properties["max_mana"]; ok {
        oComponent.SetProperty("max_mana", maxMana)
    } else {
        oComponent.SetProperty("max_mana", 50.0)
    }
    
    // Set current mana
    if currentMana, ok := p.Properties["current_mana"]; ok {
        oComponent.SetProperty("current_mana", currentMana)
    } else {
        maxMana, _ := oComponent.GetProperty("max_mana")
        oComponent.SetProperty("current_mana", maxMana)
    }
    
    // Set spell power
    if spellPower, ok := p.Properties["spell_power"]; ok {
        oComponent.SetProperty("spell_power", spellPower)
    } else {
        oComponent.SetProperty("spell_power", 10.0)
    }
    
    // Set mana regeneration
    if manaRegen, ok := p.Properties["mana_regeneration"]; ok {
        oComponent.SetProperty("mana_regeneration", manaRegen)
    } else {
        oComponent.SetProperty("mana_regeneration", 2.0)
    }
    
    // Set casting ability
    if canCast, ok := p.Flags["can_cast"]; ok {
        oComponent.SetFlag("can_cast", canCast)
    } else {
        oComponent.SetFlag("can_cast", true)
    }
    
    // Set magical school
    if school, ok := p.Fields["magical_school"]; ok {
        oComponent.SetField("magical_school", school)
    } else {
        oComponent.SetField("magical_school", "arcane")
    }
    
    return newComponent.OComponent
}

// CastSpell attempts to cast a spell, consuming mana.
func (m *Magic) CastSpell(manaCost float32) bool {
    canCast, _ := m.GetFlag("can_cast")
    if !canCast {
        return false
    }
    
    currentMana, _ := m.GetProperty("current_mana")
    if currentMana < manaCost {
        return false // Not enough mana
    }
    
    // Consume mana
    m.SetProperty("current_mana", currentMana-manaCost)
    return true
}

// RestoreMana increases the entity's current mana.
func (m *Magic) RestoreMana(amount float32) {
    current, _ := m.GetProperty("current_mana")
    max, _ := m.GetProperty("max_mana")
    
    newMana := current + amount
    if newMana > max {
        newMana = max
    }
    
    m.SetProperty("current_mana", newMana)
}
```

### 3. Container Component

Finally, a container component for items that hold other items:

```go
// File: internal/world/entity/component/components/container.go
package components

import (
    "mud-engine/internal/world/entity/component"
    "mud-engine/internal/world/entity/prototype"
)

// Container allows entities to hold other entities (bags, chests, etc.).
//
// Properties:
//   - capacity (property): Maximum weight that can be stored
//   - current_weight (property): Current weight of stored items
//   - max_items (property): Maximum number of individual items
//   - current_items (property): Current number of stored items
//   - is_open (flag): Whether the container is currently open
//   - container_type (field): Type of container (bag, chest, etc.)
type Container struct {
    component.OComponent
    contents []string // IDs of contained entities
}

// BuildComponent initializes a Container component from prototype data.
func (c *Container) BuildComponent(p *prototype.ObjectPrototype) component.OComponent {
    newComponent := Container{
        contents: make([]string, 0),
    }
    oComponent := newComponent.OComponent
    
    oComponent.Name = "Container"
    
    // Initialize maps
    if oComponent.Properties == nil {
        oComponent.Properties = make(map[string]float32)
    }
    if oComponent.Flags == nil {
        oComponent.Flags = make(map[string]bool)
    }
    if oComponent.Fields == nil {
        oComponent.Fields = make(map[string]string)
    }
    
    // Set capacity
    if capacity, ok := p.Properties["capacity"]; ok {
        oComponent.SetProperty("capacity", capacity)
    } else {
        oComponent.SetProperty("capacity", 25.0)
    }
    
    // Set maximum items
    if maxItems, ok := p.Properties["max_items"]; ok {
        oComponent.SetProperty("max_items", maxItems)
    } else {
        oComponent.SetProperty("max_items", 10.0)
    }
    
    // Initialize current values
    oComponent.SetProperty("current_weight", 0.0)
    oComponent.SetProperty("current_items", 0.0)
    
    // Set open status
    if isOpen, ok := p.Flags["is_open"]; ok {
        oComponent.SetFlag("is_open", isOpen)
    } else {
        oComponent.SetFlag("is_open", true)
    }
    
    // Set container type
    if containerType, ok := p.Fields["container_type"]; ok {
        oComponent.SetField("container_type", containerType)
    } else {
        oComponent.SetField("container_type", "generic")
    }
    
    return newComponent.OComponent
}

// AddItem attempts to add an item to the container.
func (c *Container) AddItem(itemID string, weight float32) bool {
    currentWeight, _ := c.GetProperty("current_weight")
    capacity, _ := c.GetProperty("capacity")
    currentItems, _ := c.GetProperty("current_items")
    maxItems, _ := c.GetProperty("max_items")
    
    // Check capacity constraints
    if currentWeight+weight > capacity || currentItems >= maxItems {
        return false
    }
    
    // Add the item
    c.contents = append(c.contents, itemID)
    c.SetProperty("current_weight", currentWeight+weight)
    c.SetProperty("current_items", currentItems+1)
    
    return true
}

// RemoveItem removes an item from the container.
func (c *Container) RemoveItem(itemID string, weight float32) bool {
    // Find and remove the item
    for i, id := range c.contents {
        if id == itemID {
            // Remove from slice
            c.contents = append(c.contents[:i], c.contents[i+1:]...)
            
            // Update properties
            currentWeight, _ := c.GetProperty("current_weight")
            currentItems, _ := c.GetProperty("current_items")
            c.SetProperty("current_weight", currentWeight-weight)
            c.SetProperty("current_items", currentItems-1)
            
            return true
        }
    }
    return false
}

// GetContents returns the IDs of all items in the container.
func (c *Container) GetContents() []string {
    return c.contents
}
```

## Complete Usage Example

Now let's create a complete example using these custom components:

```go
package main

import (
    "fmt"
    "mud-engine/internal/world/entity"
    "mud-engine/internal/world/entity/prototype"
    "mud-engine/internal/world/scope"
    // Note: In a real implementation, you'd import your custom components
    // "mud-engine/internal/world/entity/component/components"
)

func main() {
    fmt.Println("=== Custom Component Demonstration ===\n")
    
    // Create a test environment
    room := scope.NewScope()
    
    // Create entities with custom components
    wizard := createWizard()
    warrior := createWarrior()
    magicBag := createMagicBag()
    healthPotion := createHealthPotion()
    
    // Add entities to room
    room.AddEntity("wizard", *wizard)
    room.AddEntity("warrior", *warrior)
    room.AddEntity("magic_bag", *magicBag)
    room.AddEntity("health_potion", *healthPotion)
    
    // Demonstrate component interactions
    demonstrateHealthSystem(wizard, warrior)
    demonstrateMagicSystem(wizard)
    demonstrateContainerSystem(magicBag, healthPotion)
    
    // Show final entity states
    showEntityStates(room)
}

func createWizard() *entity.WorldObject {
    jsonStr := `{
        "name": "Gandric the Wise",
        "description": "An elderly wizard with piercing blue eyes and a long silver beard",
        "max_health": 80.0,
        "current_health": 80.0,
        "regeneration_rate": 0.5,
        "max_mana": 120.0,
        "current_mana": 120.0,
        "spell_power": 25.0,
        "mana_regeneration": 5.0,
        "can_cast": true,
        "magical_school": "evocation",
        "has_eyes": true,
        "eye_color": "blue"
    }`
    
    prototype, err := prototype.NewObjectPrototypeFromJSON(jsonStr)
    if err != nil {
        panic(fmt.Sprintf("Failed to create wizard: %v", err))
    }
    
    return entity.NewEntity("wizard", prototype)
}

func createWarrior() *entity.WorldObject {
    jsonStr := `{
        "name": "Thorek Ironshield",
        "description": "A sturdy dwarf warrior in plate armor, wielding a mighty axe",
        "max_health": 150.0,
        "current_health": 120.0,
        "regeneration_rate": 2.0,
        "death_message": "falls in glorious battle",
        "has_legs": true,
        "leg_description": "stout and powerful"
    }`
    
    prototype, err := prototype.NewObjectPrototypeFromJSON(jsonStr)
    if err != nil {
        panic(fmt.Sprintf("Failed to create warrior: %v", err))
    }
    
    return entity.NewEntity("warrior", prototype)
}

func createMagicBag() *entity.WorldObject {
    jsonStr := `{
        "name": "Enchanted Leather Bag",
        "description": "A supple leather bag that seems larger inside than outside",
        "capacity": 50.0,
        "max_items": 15.0,
        "is_open": true,
        "container_type": "bag",
        "magical": true
    }`
    
    prototype, err := prototype.NewObjectPrototypeFromJSON(jsonStr)
    if err != nil {
        panic(fmt.Sprintf("Failed to create bag: %v", err))
    }
    
    return entity.NewEntity("magic_bag", prototype)
}

func createHealthPotion() *entity.WorldObject {
    jsonStr := `{
        "name": "Healing Potion",
        "description": "A small vial filled with red liquid that glows softly",
        "weight": 0.5,
        "healing_power": 30.0,
        "potion_type": "healing"
    }`
    
    prototype, err := prototype.NewObjectPrototypeFromJSON(jsonStr)
    if err != nil {
        panic(fmt.Sprintf("Failed to create potion: %v", err))
    }
    
    return entity.NewEntity("health_potion", prototype)
}

func demonstrateHealthSystem(wizard, warrior *entity.WorldObject) {
    fmt.Println("⚔️  === Health System Demo ===")
    
    // Note: In a real implementation, you'd access the Health component
    // and call its methods. For this example, we'll simulate the behavior.
    
    fmt.Printf("Wizard starts with 80/80 HP\n")
    fmt.Printf("Warrior starts with 120/150 HP\n")
    
    fmt.Println("Wizard casts fireball at warrior...")
    fmt.Printf("Warrior takes 25 damage (120 -> 95 HP)\n")
    
    fmt.Println("Warrior regenerates 2 HP per turn...")
    fmt.Printf("Warrior heals (95 -> 97 HP)\n")
    
    fmt.Println()
}

func demonstrateMagicSystem(wizard *entity.WorldObject) {
    fmt.Println("✨ === Magic System Demo ===")
    
    fmt.Printf("Wizard starts with 120/120 mana\n")
    
    fmt.Println("Wizard casts Lightning Bolt (costs 15 mana)...")
    fmt.Printf("Mana: 120 -> 105\n")
    
    fmt.Println("Wizard casts Fireball (costs 20 mana)...")
    fmt.Printf("Mana: 105 -> 85\n")
    
    fmt.Println("Wizard regenerates 5 mana per turn...")
    fmt.Printf("Mana: 85 -> 90\n")
    
    fmt.Println()
}

func demonstrateContainerSystem(bag, potion *entity.WorldObject) {
    fmt.Println("🎒 === Container System Demo ===")
    
    fmt.Printf("Magic bag capacity: 50.0 weight, 15 items\n")
    fmt.Printf("Current contents: 0.0 weight, 0 items\n")
    
    fmt.Println("Adding healing potion (weight: 0.5)...")
    fmt.Printf("Bag contents: 0.5 weight, 1 item\n")
    
    fmt.Println("Bag contents: [health_potion]\n")
    
    fmt.Println()
}

func showEntityStates(room *scope.Scope) {
    fmt.Println("📊 === Final Entity States ===")
    
    entities := []string{"wizard", "warrior", "magic_bag", "health_potion"}
    
    for _, entityID := range entities {
        if entity, exists := room.GetEntity(entityID); exists {
            fmt.Printf("\n%s (%s):\n", entity.Name(), entityID)
            
            components := entity.Components()
            for _, comp := range components {
                fmt.Printf("  %s Component:\n", comp.Name)
                
                // Show component data
                for key, value := range comp.Properties {
                    if value > 0 {
                        fmt.Printf("    %s: %.1f\n", key, value)
                    }
                }
                
                for key, value := range comp.Flags {
                    if value {
                        fmt.Printf("    %s: enabled\n", key)
                    }
                }
                
                for key, value := range comp.Fields {
                    if value != "" {
                        fmt.Printf("    %s: %s\n", key, value)
                    }
                }
            }
        }
    }
}
```

## Key Concepts Demonstrated

### 1. Component Structure
- Embed `component.OComponent` as base
- Initialize all maps in `BuildComponent`
- Provide sensible default values
- Implement component-specific methods

### 2. Data Organization
- **Properties**: Numeric values (health, mana, capacity)
- **Flags**: Boolean states (alive, can_cast, is_open)
- **Fields**: String data (descriptions, types, schools)

### 3. Component Interaction
- Components can interact through shared entity
- Cross-component logic (magic healing, container weight)
- Event-driven updates between components

### 4. Extension Points
- Add new component methods for specific behavior
- Implement component interfaces for system integration  
- Create component dependencies and interactions

## Running the Example

Save the component code in appropriate files under:
```
internal/world/entity/component/components/
├── health.go
├── magic.go
└── container.go
```

Then run the main example:
```bash
go run custom_component_example.go
```

## Benefits of Custom Components

1. **Modularity**: Each component handles one concern
2. **Reusability**: Components work on any entity
3. **Flexibility**: Mix and match components as needed
4. **Maintainability**: Isolated, testable code units
5. **Extensibility**: Easy to add new capabilities

## Related Examples

- [Basic Entity Creation](basic-entity-creation.md) - Entity fundamentals
- [Simple Room](simple-room.md) - Complete room with interactions
- [Component Examples](component-examples.md) - Using built-in components