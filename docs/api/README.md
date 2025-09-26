# WikiMUD Engine API Reference

## Core Types

### WorldObject (Entity)

The fundamental game object representing any entity in the MUD world.

```go
type WorldObject struct {
    id          string
    name        string
    description string
    components  *ComponentManager
    namespace   string
    parent      *WorldObject
}
```

#### Methods

##### `NewEntity(id, prototype *prototype.ObjectPrototype) *WorldObject`
Creates a new world object from a prototype.

**Parameters:**
- `id`: Unique identifier for the entity
- `prototype`: Object prototype defining initial properties

**Returns:** Pointer to newly created WorldObject

##### `ID() string`
Returns the unique identifier of the entity.

##### `Name() string`
Returns the display name of the entity.

##### `Description() string` 
Returns the detailed description of the entity.

##### `Components() []component.OComponent`
Returns all components attached to this entity.

##### `Namespace() string`
Returns the namespace this entity belongs to.

##### `Parent() *WorldObject`
Returns the parent entity (for containment hierarchy).

##### `EmitEvent(event string, params map[string]interface{})`
Emits an event from this entity.

**Parameters:**
- `event`: Event type identifier
- `params`: Event-specific parameters

##### `HandleEvent(event string, params map[string]interface{})`
Handles an incoming event. Override this method for custom event processing.

### Component System

#### OComponent

Base component structure that all components inherit from.

```go
type OComponent struct {
    Name       string
    Flags      map[string]bool
    Fields     map[string]string
    Properties map[string]float32
    Initialize func(c *OComponent, o *prototype.ObjectPrototype) error
    Refresh    func(c *OComponent) error
}
```

#### Methods

##### `SetProperty(key string, value float32)`
Sets a numeric property on the component.

##### `GetProperty(key string) (float32, bool)`
Retrieves a numeric property. Returns value and existence flag.

##### `SetField(key string, value string)`
Sets a string field on the component.

##### `GetField(key string) (string, bool)`
Retrieves a string field. Returns value and existence flag.

##### `SetFlag(key string, value bool)`
Sets a boolean flag on the component.

##### `GetFlag(key string) (bool, bool)`
Retrieves a boolean flag. Returns value and existence flag.

##### `BuildComponent(proto *ObjectPrototype) *OComponent`
Builds component from prototype data. Override in specific components.

#### ComponentManager

Manages the lifecycle and collection of components for an entity.

```go
type ComponentManager struct {
    Components []OComponent
}
```

##### `Initialize(prototype *ObjectPrototype)`
Initializes all components from a prototype.

##### `GetComponents() []OComponent`
Returns all components managed by this manager.

##### `NewComponentManager(prototype *ObjectPrototype) *ComponentManager`
Creates a new component manager initialized with prototype data.

### Built-in Components

#### HasEyes

Manages vision and sight capabilities.

```go
type HasEyes struct {
    component.OComponent
}
```

**Component Properties:**
- `has_eyes` (flag): Whether entity has functional eyes
- `eye_color` (field): Color of the eyes
- `eye_count` (property): Number of eyes

##### `BuildComponent(p *ObjectPrototype) component.OComponent`
Builds HasEyes component from prototype data.

**Default Values:**
- `has_eyes`: false
- `eye_color`: "" (empty)
- `eye_count`: 2

#### HasLegs

Manages movement and locomotion capabilities.

```go
type HasLegs struct {
    component.OComponent
}
```

**Component Properties:**
- `leg_description` (field): Description of legs
- `leg_count` (property): Number of legs

##### `BuildComponent(p *ObjectPrototype) component.OComponent`
Builds HasLegs component from prototype data.

**Default Values:**
- `leg_description`: "very leglike"
- `leg_count`: 2

#### CanOpen

Manages interaction with openable objects (doors, containers).

```go
type CanOpen struct {
    component.OComponent
}
```

**Component Properties:**
- `can_open` (flag): Can this entity open things
- `can_lock` (flag): Can this entity be locked
- `is_locked` (flag): Current lock state
- `key` (field): Required key identifier

##### `BuildComponent(p *ObjectPrototype) component.OComponent`
Builds CanOpen component from prototype data.

**Default Values:**
- `can_open`: false
- `can_lock`: false
- `is_locked`: true (if can_lock is true)
- `key`: "" (empty)

### Prototype System

#### ObjectPrototype

Defines object templates for entity creation.

```go
type ObjectPrototype struct {
    Fields     map[string]string
    Flags      map[string]bool
    Properties map[string]float32
}
```

##### `NewObjectPrototypeFromJSON(jsonStr string) (*ObjectPrototype, error)`
Creates a prototype from JSON string.

**Parameters:**
- `jsonStr`: JSON representation of the prototype

**Returns:** Prototype object and error (if any)

**JSON Format:**
```json
{
    "string_field": "string_value",
    "boolean_flag": true,
    "numeric_property": 42.5
}
```

### World Management

#### GameWorld

Top-level container for the entire game world.

```go
type GameWorld struct {
    objects      map[string]entity.WorldObject
    verbs        []verb.VerbType
    scopes       map[string]interface{}
    eventHandler *event.EventHandler
}
```

##### `add_object(id string, parameters prototype.ObjectPrototype)`
Adds an object to the world.

**Parameters:**
- `id`: Unique identifier for the object
- `parameters`: Prototype defining object properties

##### `object_by_id(id string) (entity.WorldObject, bool)`
Retrieves an object by ID.

**Returns:** Object and existence flag

##### `register_verb(verb verb.VerbType)`
Registers a new verb with the world.

#### Scope

Manages entities within a specific context (room, area, etc.).

```go
type Scope struct {
    entities map[string]entity.WorldObject
}
```

##### `NewScope() *Scope`
Creates a new empty scope.

##### `AddEntity(id string, obj entity.WorldObject)`
Adds an entity to this scope.

##### `RemoveEntity(id string)`
Removes an entity from this scope.

##### `GetEntity(id string) (entity.WorldObject, bool)`
Retrieves an entity by ID. Returns entity and existence flag.

### Verb System

#### Verb

Represents a player action or command.

```go
type Verb struct {
    Name    string
    Actor   entity.WorldObject
    Targets []entity.WorldObject
    Params  []string
    Type    VerbType
}
```

##### `Try(actor entity.WorldObject, verbType VerbType, targets []entity.WorldObject, params []string) string`
Executes the verb and returns result description.

**Parameters:**
- `actor`: Entity performing the action
- `verbType`: Type of verb being executed
- `targets`: Target entities for the action
- `params`: Additional parameters

**Returns:** String description of the action result

#### VerbType

Defines a type of verb that can be executed.

```go
type VerbType struct {
    Name string
}
```

### Rendering System

#### Renderer

Base renderer for converting game state to output.

```go
type Renderer struct {
}
```

##### `Render(scope scope.Scope) string`
Renders a scope to string output.

**Parameters:**
- `scope`: The scope to render

**Returns:** Formatted string representation

### Event System

#### EventHandler

Manages event routing and processing.

```go
type EventHandler struct {
    // Implementation details
}
```

Events flow through the system as string identifiers with parameter maps. The system supports:

- Event emission from any WorldObject
- Event handling by registered handlers
- Parameter passing through interface{} maps
- Asynchronous event processing (planned)

## Usage Examples

### Creating an Entity

```go
// Define prototype
jsonStr := `{
    "name": "magic_sword",
    "description": "A gleaming sword",
    "damage": 15.0,
    "magical": true
}`

prototype, err := prototype.NewObjectPrototypeFromJSON(jsonStr)
if err != nil {
    log.Fatal(err)
}

// Create entity
entity := entity.NewEntity("sword_001", prototype)
```

### Adding Components

```go
// Components are automatically created from prototypes
// when the entity is initialized

// Access component data
if hasEyes := entity.Components().HasEyes; hasEyes != nil {
    eyeColor, exists := hasEyes.GetField("eye_color")
    eyeCount, exists := hasEyes.GetProperty("eye_count")
}
```

### Working with Scopes

```go
// Create a room scope
room := scope.NewScope()
room.AddEntity("player_001", player)
room.AddEntity("sword_001", sword)

// Render the room
renderer := render.Renderer{}
output := renderer.Render(*room)
```

### Executing Verbs

```go
// Create and execute a verb
verb := verb.Verb{}
result := verb.Try(player, lookVerb, []entity.WorldObject{sword}, []string{})
fmt.Println(result) // "look with target: magic_sword"
```

## Error Handling

The API uses Go's standard error handling patterns:

- Functions that can fail return `(result, error)` tuples
- Check error values before using results
- Components use boolean flags to indicate property existence
- Invalid operations return safe default values where appropriate

## Thread Safety

Current implementation is not thread-safe. For concurrent access:

- Protect shared data structures with mutexes
- Use channels for entity communication
- Consider read-write locks for performance
- Implement atomic operations for counters