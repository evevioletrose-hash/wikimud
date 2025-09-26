# Creating Entities - Basic Example

This example shows how to create entities (game objects) using WikiMUD Engine's prototype system.

## Overview

We'll create several different types of entities:
- A player character with eyes and legs
- A magical sword that can be opened (has a hidden compartment)
- A locked chest that requires a key

## Complete Example

```go
package main

import (
    "fmt"
    "mud-engine/internal/world/entity"
    "mud-engine/internal/world/entity/prototype"
    "mud-engine/internal/world/scope"
)

func main() {
    // Create a scope to hold our entities (like a room)
    room := scope.NewScope()
    
    // Example 1: Create a player character
    playerPrototype := createPlayerPrototype()
    player := entity.NewEntity("player_001", playerPrototype)
    room.AddEntity("player_001", *player)
    
    // Example 2: Create a magical sword with hidden compartment
    swordPrototype := createSwordPrototype()
    sword := entity.NewEntity("magic_sword", swordPrototype)
    room.AddEntity("magic_sword", *sword)
    
    // Example 3: Create a locked treasure chest
    chestPrototype := createChestPrototype()
    chest := entity.NewEntity("treasure_chest", chestPrototype)
    room.AddEntity("treasure_chest", *chest)
    
    // Display information about our entities
    displayEntityInfo(player, "Player Character")
    displayEntityInfo(sword, "Magic Sword")
    displayEntityInfo(chest, "Treasure Chest")
    
    // Show scope contents
    fmt.Println("\n=== Room Contents ===")
    if playerInRoom, exists := room.GetEntity("player_001"); exists {
        fmt.Printf("Player: %s\n", playerInRoom.Name())
    }
    if swordInRoom, exists := room.GetEntity("magic_sword"); exists {
        fmt.Printf("Sword: %s\n", swordInRoom.Name())
    }
    if chestInRoom, exists := room.GetEntity("treasure_chest"); exists {
        fmt.Printf("Chest: %s\n", chestInRoom.Name())
    }
}

// createPlayerPrototype creates a prototype for a player character
func createPlayerPrototype() *prototype.ObjectPrototype {
    jsonStr := `{
        "name": "Brave Adventurer",
        "description": "A courageous explorer ready for adventure",
        "has_eyes": true,
        "eye_color": "green",
        "eye_count": 2,
        "leg_description": "strong and sturdy",
        "leg_count": 2
    }`
    
    proto, err := prototype.NewObjectPrototypeFromJSON(jsonStr)
    if err != nil {
        panic(fmt.Sprintf("Failed to create player prototype: %v", err))
    }
    return proto
}

// createSwordPrototype creates a magical sword with a hidden compartment
func createSwordPrototype() *prototype.ObjectPrototype {
    jsonStr := `{
        "name": "Gleaming Blade of Power",
        "description": "A magnificently crafted sword that seems to hum with magical energy",
        "can_open": true,
        "can_lock": false,
        "is_locked": false,
        "damage": 25.5,
        "magical": true,
        "weight": 3.2
    }`
    
    proto, err := prototype.NewObjectPrototypeFromJSON(jsonStr)
    if err != nil {
        panic(fmt.Sprintf("Failed to create sword prototype: %v", err))
    }
    return proto
}

// createChestPrototype creates a locked treasure chest
func createChestPrototype() *prototype.ObjectPrototype {
    jsonStr := `{
        "name": "Ancient Treasure Chest",
        "description": "An ornate wooden chest bound with iron, clearly very old",
        "can_open": true,
        "can_lock": true,
        "is_locked": true,
        "key": "brass_key_001",
        "capacity": 50.0,
        "durability": 85.5
    }`
    
    proto, err := prototype.NewObjectPrototypeFromJSON(jsonStr)
    if err != nil {
        panic(fmt.Sprintf("Failed to create chest prototype: %v", err))
    }
    return proto
}

// displayEntityInfo shows detailed information about an entity
func displayEntityInfo(entity *entity.WorldObject, title string) {
    fmt.Printf("\n=== %s ===\n", title)
    fmt.Printf("ID: %s\n", entity.ID())
    fmt.Printf("Name: %s\n", entity.Name())
    fmt.Printf("Description: %s\n", entity.Description())
    
    fmt.Println("Components:")
    components := entity.Components()
    for i, comp := range components {
        fmt.Printf("  Component %d: %s\n", i+1, comp.Name)
        
        // Show component fields
        if len(comp.Fields) > 0 {
            fmt.Println("    Fields:")
            for key, value := range comp.Fields {
                fmt.Printf("      %s: %s\n", key, value)
            }
        }
        
        // Show component flags
        if len(comp.Flags) > 0 {
            fmt.Println("    Flags:")
            for key, value := range comp.Flags {
                fmt.Printf("      %s: %t\n", key, value)
            }
        }
        
        // Show component properties
        if len(comp.Properties) > 0 {
            fmt.Println("    Properties:")
            for key, value := range comp.Properties {
                fmt.Printf("      %s: %.2f\n", key, value)
            }
        }
    }
}
```

## Step-by-Step Explanation

### 1. Define Prototypes with JSON

Prototypes are defined using JSON strings that specify the initial state of entities:

```go
jsonStr := `{
    "name": "Brave Adventurer",
    "description": "A courageous explorer ready for adventure",
    "has_eyes": true,
    "eye_color": "green",
    "eye_count": 2
}`
```

The JSON fields are automatically categorized by type:
- **String values** → Component Fields
- **Boolean values** → Component Flags  
- **Numeric values** → Component Properties

### 2. Create Prototype Objects

Convert JSON strings to prototype objects:

```go
proto, err := prototype.NewObjectPrototypeFromJSON(jsonStr)
if err != nil {
    // Handle error
}
```

### 3. Create Entities from Prototypes

Use prototypes to create actual game entities:

```go
entity := entity.NewEntity("unique_id", prototype)
```

Each entity gets:
- A unique identifier
- Properties from the prototype
- Components automatically created based on prototype data

### 4. Add Entities to Scopes

Entities exist within scopes (rooms, areas, etc.):

```go
scope := scope.NewScope()
scope.AddEntity("entity_id", entity)
```

### 5. Access Entity Data

Retrieve entities and their properties:

```go
if entity, exists := scope.GetEntity("entity_id"); exists {
    fmt.Println("Name:", entity.Name())
    fmt.Println("Description:", entity.Description())
}
```

## Component System Integration

The WikiMUD engine automatically creates components based on prototype data:

### HasEyes Component
Activated by:
- `has_eyes` (boolean) → Flag
- `eye_color` (string) → Field  
- `eye_count` (number) → Property

### HasLegs Component
Activated by:
- `leg_description` (string) → Field
- `leg_count` (number) → Property

### CanOpen Component
Activated by:
- `can_open` (boolean) → Flag
- `can_lock` (boolean) → Flag
- `is_locked` (boolean) → Flag
- `key` (string) → Field

## Running the Example

Save the code as `entity_example.go` and run:

```bash
go run entity_example.go
```

Expected output shows detailed information for each created entity, including their components and properties.

## Extensions

Try these modifications:

1. **Add More Entities**: Create NPCs, monsters, or special items
2. **Custom Properties**: Add health, mana, experience points
3. **Relationships**: Create parent-child entity relationships
4. **Collections**: Store related entities together

## Related Examples

- [Component Usage](component-examples.md) - Working directly with components
- [Prototype Definitions](prototype-examples.md) - Advanced prototype patterns
- [Simple Room](simple-room.md) - Building a complete game room