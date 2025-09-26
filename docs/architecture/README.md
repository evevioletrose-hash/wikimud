# WikiMUD Engine Architecture

## Overview

WikiMUD is built on a modern Entity-Component-System (ECS) architecture that provides flexibility, modularity, and extensibility for creating text-based multiplayer games.

## Core Concepts

### Entity-Component-System (ECS)

The ECS pattern separates data from behavior, allowing for highly flexible object composition:

- **Entities**: Unique identifiers for game objects (players, items, rooms, NPCs)
- **Components**: Data containers that define specific aspects of entities
- **Systems**: Logic that operates on entities with specific component combinations

```
Entity (WorldObject)
├── ID: "player_123"
├── Name: "Brave Adventurer"
├── Description: "A courageous explorer"
└── Components:
    ├── HasEyes (component)
    │   ├── eye_color: "blue"
    │   ├── eye_count: 2
    │   └── has_eyes: true
    ├── HasLegs (component)
    │   ├── leg_description: "strong"
    │   └── leg_count: 2
    └── CanOpen (component)
        ├── can_open: true
        ├── can_lock: false
        └── key: ""
```

### World Hierarchy

```
GameWorld
├── Objects: map[string]WorldObject
├── Verbs: []VerbType
├── Scopes: map[string]Scope
└── EventHandler: *EventHandler

Scope (Room/Area)
├── Entities: map[string]WorldObject
├── AddEntity(id, object)
├── RemoveEntity(id)
└── GetEntity(id)

WorldObject (Entity)
├── ID: string
├── Name: string
├── Description: string
├── Components: *ComponentManager
├── Namespace: string
└── Parent: *WorldObject
```

## Component System

### Component Structure

Each component inherits from `OComponent` and provides:

```go
type OComponent struct {
    Name       string
    Flags      map[string]bool       // Boolean properties
    Fields     map[string]string     // String properties
    Properties map[string]float32    // Numeric properties
    Initialize func(c *OComponent, o *prototype.ObjectPrototype) error
    Refresh    func(c *OComponent) error
}
```

### Built-in Components

#### HasEyes Component
Manages vision and sight mechanics:
- `has_eyes` (flag): Whether the entity can see
- `eye_color` (field): Color of the eyes
- `eye_count` (property): Number of eyes

#### HasLegs Component
Handles movement and locomotion:
- `leg_description` (field): Description of legs
- `leg_count` (property): Number of legs

#### CanOpen Component
Controls interaction with containers and doors:
- `can_open` (flag): Can this entity open things
- `can_lock` (flag): Can this entity lock things
- `is_locked` (flag): Current lock state
- `key` (field): Required key identifier

### Component Manager

The `ComponentManager` orchestrates component lifecycle:

```go
type ComponentManager struct {
    Components []OComponent
}

func (cm *ComponentManager) Initialize(prototype *ObjectPrototype)
func (cm *ComponentManager) GetComponents() []OComponent
```

## Prototype System

Prototypes define object templates using JSON configuration:

```go
type ObjectPrototype struct {
    Fields     map[string]string    // String values
    Flags      map[string]bool      // Boolean values
    Properties map[string]float32   // Numeric values
}
```

Example JSON prototype:
```json
{
    "name": "wooden_door",
    "description": "A sturdy wooden door",
    "can_open": true,
    "can_lock": true,
    "is_locked": false,
    "key": "brass_key",
    "durability": 100.0
}
```

## Verb System

Verbs represent player actions and commands:

```go
type Verb struct {
    Name    string
    Actor   WorldObject      // Who performs the action
    Targets []WorldObject    // What the action targets
    Params  []string         // Additional parameters
    Type    VerbType         // Verb classification
}
```

### Verb Execution Flow

1. Parse player input into verb + targets + parameters
2. Validate targets and permissions
3. Execute verb logic based on target count:
   - 0 targets: Intransitive action (e.g., "look")
   - 1 target: Transitive action (e.g., "get sword")
   - Multiple targets: Complex action (e.g., "put sword in bag")

## Rendering System

The rendering system converts game state into player-readable output:

### Renderer Types

#### Raw Text Renderer
Basic plain text output for traditional MUD experience.

#### Linked Text Renderer
Enhanced text with clickable elements for modern interfaces.

#### Journal View Renderer
Structured narrative format for story-driven experiences.

### Rendering Flow

```
Scope → Renderer → Output
  ↓
Entities → Components → Data → Format → Player
```

## Event System

Events decouple game logic and enable reactive programming:

```go
type EventHandler struct {
    // Event routing and handling logic
}

func (wo *WorldObject) EmitEvent(event string, params map[string]interface{})
func (wo *WorldObject) HandleEvent(event string, params map[string]interface{})
```

### Event Types

- **Movement Events**: Entity position changes
- **Interaction Events**: Object interactions
- **Combat Events**: Fighting and damage
- **Communication Events**: Chat and messaging
- **System Events**: Login, logout, errors

## Data Flow

```
Player Input → Verb Parser → Action Validation → Component Logic → 
State Change → Event Emission → Scope Updates → Rendering → Output
```

## Extensibility Points

### Adding New Components

1. Create struct embedding `OComponent`
2. Implement `BuildComponent` method
3. Register with component system
4. Define prototype fields/flags/properties

### Adding New Verbs

1. Define `VerbType` structure
2. Implement verb logic in `Try` method
3. Register with world verb system
4. Handle target validation

### Adding New Renderers

1. Implement `Renderer` interface
2. Handle scope-to-output conversion
3. Register with rendering system
4. Support specific output formats

## Performance Considerations

- Components use maps for O(1) property access
- Entity lookup uses string-based indexing
- Scope isolation prevents unnecessary updates
- Event system enables lazy evaluation
- Component composition avoids inheritance overhead

## Thread Safety

Current implementation is single-threaded. For multiplayer support:
- Add mutex protection to shared data structures
- Implement goroutine-safe event dispatching
- Use channels for inter-entity communication
- Consider actor model for entity isolation