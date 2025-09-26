# Simple Room Example

This example creates a complete MUD room with a player, items, and demonstrates basic interactions using the WikiMUD Engine.

## Overview

We'll create:
- A tavern room with atmosphere
- A player character
- Interactive items (a glowing orb, a locked chest)
- Basic verb interactions (look, examine)
- Multiple rendering formats

## Complete Example

Save this as `simple_room_example.go`:

```go
package main

import (
    "fmt"
    "strings"
    "mud-engine/internal/world/entity"
    "mud-engine/internal/world/entity/prototype"
    "mud-engine/internal/world/entity/verb"
    "mud-engine/internal/world/scope"
    "mud-engine/internal/world/render"
)

func main() {
    fmt.Println("=== WikiMUD Engine: Simple Room Example ===\n")
    
    // Create the tavern room scope
    tavern := createTavernRoom()
    
    // Create and add entities to the room
    setupRoomEntities(tavern)
    
    // Demonstrate different rendering methods
    demonstrateRendering(tavern)
    
    // Demonstrate verb interactions
    demonstrateVerbs(tavern)
    
    // Show final room state
    showRoomContents(tavern)
}

// createTavernRoom sets up the basic room scope
func createTavernRoom() *scope.Scope {
    room := scope.NewScope()
    fmt.Println("🏛️  Created 'The Prancing Pony' tavern room")
    return room
}

// setupRoomEntities creates and adds all entities to the room
func setupRoomEntities(room *scope.Scope) {
    // Create player character
    player := createPlayer()
    room.AddEntity("player", *player)
    fmt.Printf("👤 Added player: %s\n", player.Name())
    
    // Create a magical orb
    orb := createMagicalOrb()
    room.AddEntity("glowing_orb", *orb)
    fmt.Printf("✨ Added item: %s\n", orb.Name())
    
    // Create a locked treasure chest
    chest := createTreasureChest()
    room.AddEntity("treasure_chest", *chest)
    fmt.Printf("📦 Added item: %s\n", chest.Name())
    
    // Create a brass key
    key := createBrassKey()
    room.AddEntity("brass_key", *key)
    fmt.Printf("🗝️  Added item: %s\n", key.Name())
    
    fmt.Println()
}

// createPlayer creates the player character
func createPlayer() *entity.WorldObject {
    jsonStr := `{
        "name": "Elara the Explorer",
        "description": "A keen-eyed adventurer with an insatiable curiosity for ancient mysteries",
        "has_eyes": true,
        "eye_color": "emerald green",
        "eye_count": 2,
        "leg_description": "long and athletic",
        "leg_count": 2
    }`
    
    prototype, err := prototype.NewObjectPrototypeFromJSON(jsonStr)
    if err != nil {
        panic(fmt.Sprintf("Failed to create player: %v", err))
    }
    
    return entity.NewEntity("player", prototype)
}

// createMagicalOrb creates a glowing magical orb
func createMagicalOrb() *entity.WorldObject {
    jsonStr := `{
        "name": "Glowing Crystal Orb",
        "description": "A perfectly spherical crystal that pulses with inner light, warm to the touch and humming with arcane energy",
        "magical": true,
        "luminosity": 3.5,
        "temperature": 98.6,
        "can_open": false
    }`
    
    prototype, err := prototype.NewObjectPrototypeFromJSON(jsonStr)
    if err != nil {
        panic(fmt.Sprintf("Failed to create orb: %v", err))
    }
    
    return entity.NewEntity("glowing_orb", prototype)
}

// createTreasureChest creates a locked chest that can be opened with a key
func createTreasureChest() *entity.WorldObject {
    jsonStr := `{
        "name": "Ornate Iron Chest",
        "description": "A heavy chest made of dark iron with intricate engravings of dragons and ancient runes",
        "can_open": true,
        "can_lock": true,
        "is_locked": true,
        "key": "brass_key",
        "weight": 45.5,
        "capacity": 100.0
    }`
    
    prototype, err := prototype.NewObjectPrototypeFromJSON(jsonStr)
    if err != nil {
        panic(fmt.Sprintf("Failed to create chest: %v", err))
    }
    
    return entity.NewEntity("treasure_chest", prototype)
}

// createBrassKey creates the key that opens the chest
func createBrassKey() *entity.WorldObject {
    jsonStr := `{
        "name": "Tarnished Brass Key",
        "description": "An old brass key with an elaborate bow, tarnished with age but still functional",
        "can_open": false,
        "weight": 0.2,
        "material": "brass"
    }`
    
    prototype, err := prototype.NewObjectPrototypeFromJSON(jsonStr)
    if err != nil {
        panic(fmt.Sprintf("Failed to create key: %v", err))
    }
    
    return entity.NewEntity("brass_key", prototype)
}

// demonstrateRendering shows different rendering approaches
func demonstrateRendering(room *scope.Scope) {
    fmt.Println("🎨 === Rendering Demonstration ===")
    
    // Basic room description
    fmt.Println("\n📍 Room Overview:")
    fmt.Println("The Prancing Pony - A cozy tavern room")
    fmt.Println("Warm firelight flickers across worn wooden tables and stone walls.")
    fmt.Println("The air smells of ale, roasted meat, and old parchment.")
    
    // Use the built-in renderer
    renderer := &render.Renderer{}
    rendered := renderer.Render(*room)
    fmt.Printf("\nRendered output: %s\n", rendered)
    
    fmt.Println()
}

// demonstrateVerbs shows verb interactions with entities
func demonstrateVerbs(room *scope.Scope) {
    fmt.Println("⚔️  === Verb Demonstrations ===")
    
    // Get entities for verb demonstrations
    player, _ := room.GetEntity("player")
    orb, _ := room.GetEntity("glowing_orb")
    chest, _ := room.GetEntity("treasure_chest")
    key, _ := room.GetEntity("brass_key")
    
    // Create verb instances
    lookVerb := verb.VerbType{Name: "look"}
    examineVerb := verb.VerbType{Name: "examine"}
    
    // Demonstrate looking around (no target)
    lookAction := &verb.Verb{}
    result := lookAction.Try(player, lookVerb, []entity.WorldObject{}, []string{})
    fmt.Printf("👀 Player action: %s\n", result)
    
    // Demonstrate examining the orb
    examineAction := &verb.Verb{}
    result = examineAction.Try(player, examineVerb, []entity.WorldObject{orb}, []string{"closely"})
    fmt.Printf("🔍 Player action: %s\n", result)
    
    // Demonstrate examining multiple objects
    multiExamine := &verb.Verb{}
    result = multiExamine.Try(player, examineVerb, []entity.WorldObject{chest, key}, []string{"for", "clues"})
    fmt.Printf("🔍 Player action: %s\n", result)
    
    fmt.Println()
}

// showRoomContents displays detailed information about all entities in the room
func showRoomContents(room *scope.Scope) {
    fmt.Println("📋 === Detailed Room Contents ===")
    
    entities := []string{"player", "glowing_orb", "treasure_chest", "brass_key"}
    
    for _, entityID := range entities {
        if entity, exists := room.GetEntity(entityID); exists {
            showEntityDetails(entity, entityID)
        }
    }
}

// showEntityDetails displays comprehensive information about an entity
func showEntityDetails(entity entity.WorldObject, id string) {
    fmt.Printf("\n--- %s (%s) ---\n", entity.Name(), id)
    fmt.Printf("Description: %s\n", entity.Description())
    
    components := entity.Components()
    if len(components) > 0 {
        fmt.Println("Components:")
        for _, comp := range components {
            fmt.Printf("  • %s\n", comp.Name)
            
            // Show interesting component data
            showComponentData(comp)
        }
    } else {
        fmt.Println("  No components attached")
    }
}

// showComponentData displays relevant component properties
func showComponentData(comp component.OComponent) {
    var details []string
    
    // Show flags
    for key, value := range comp.Flags {
        if value {
            details = append(details, fmt.Sprintf("%s: enabled", key))
        }
    }
    
    // Show notable fields
    for key, value := range comp.Fields {
        if value != "" {
            details = append(details, fmt.Sprintf("%s: %s", key, value))
        }
    }
    
    // Show properties with values > 0
    for key, value := range comp.Properties {
        if value > 0 {
            details = append(details, fmt.Sprintf("%s: %.1f", key, value))
        }
    }
    
    if len(details) > 0 {
        fmt.Printf("    %s\n", strings.Join(details, ", "))
    }
}
```

## Step-by-Step Explanation

### 1. Room Setup

```go
room := scope.NewScope()
```

Create a scope to contain all entities. In a full MUD, this would represent a room, area, or zone.

### 2. Entity Creation Pattern

Each entity follows the same pattern:
1. Define JSON prototype with properties
2. Convert to ObjectPrototype
3. Create entity with unique ID
4. Add to room scope

```go
jsonStr := `{"name": "Entity Name", "property": value}`
prototype, err := prototype.NewObjectPrototypeFromJSON(jsonStr)
entity := entity.NewEntity("unique_id", prototype)
room.AddEntity("unique_id", *entity)
```

### 3. Component Activation

Components are automatically activated based on prototype properties:

- **HasEyes**: `has_eyes`, `eye_color`, `eye_count`
- **HasLegs**: `leg_description`, `leg_count`  
- **CanOpen**: `can_open`, `can_lock`, `is_locked`, `key`

### 4. Verb System

Verbs represent player actions:

```go
verb := verb.VerbType{Name: "look"}
action := &verb.Verb{}
result := action.Try(actor, verb, targets, params)
```

The result varies based on target count:
- 0 targets: "look with no targets"
- 1 target: "look with target: Glowing Crystal Orb"
- Multiple: "look with multiple targets: chest, key"

### 5. Rendering

The render system converts scope data to formatted output:

```go
renderer := &render.Renderer{}
output := renderer.Render(*room)
```

## Running the Example

```bash
# Save the code as simple_room_example.go
go run simple_room_example.go
```

Expected output shows:
- Room creation messages
- Entity addition confirmations  
- Rendered room description
- Verb interaction results
- Detailed entity information with components

## Key Features Demonstrated

### Entity-Component System
- Entities with multiple components
- Component data (flags, fields, properties)
- Automatic component activation from prototypes

### Prototype System
- JSON-based entity definitions
- Type inference (string → field, bool → flag, number → property)
- Flexible property specification

### Scope Management
- Entity containment and organization
- ID-based entity lookup
- Room-like entity grouping

### Verb System
- Action representation and execution
- Target validation and handling
- Parameter passing to actions

### Rendering
- Scope-to-text conversion
- Multiple rendering strategies
- Formatted output generation

## Extensions

Try these modifications:

1. **Add More Entities**: Books, furniture, NPCs
2. **Custom Components**: Health, magic, inventory systems  
3. **Advanced Verbs**: Get, drop, unlock, cast spells
4. **Room Connections**: Link multiple rooms together
5. **Interactive NPCs**: Entities that respond to events
6. **Container System**: Items that hold other items

## Related Examples

- [Basic Entity Creation](basic-entity-creation.md) - Entity creation fundamentals
- [Component Examples](component-examples.md) - Advanced component usage
- [Custom Component](custom-component.md) - Building your own components