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

#### HasHands

Enables Take, Hold, Move, Open, Close, Drop, Pickpocket verbs.

```go
type HasHands struct {
    component.OComponent
}
```

**Component Properties:**
- `has_hands` (flag): Whether entity has functional hands
- `hand_count` (property): Number of hands
- `hand_description` (field): Description of hands

**Default Values:**
- `has_hands`: false
- `hand_count`: 2
- `hand_description`: "dexterous hands"

#### HasEars

Enables Listen verb and receive heard events in scope.

```go
type HasEars struct {
    component.OComponent
}
```

**Component Properties:**
- `has_ears` (flag): Whether entity has functional ears
- `hearing_acuity` (property): Hearing sensitivity
- `ear_description` (field): Description of ears

**Default Values:**
- `has_ears`: false
- `hearing_acuity`: 1.0
- `ear_description`: "keen ears"

#### HasMouth

Enables Say, Eat, Drink, Taste verbs.

```go
type HasMouth struct {
    component.OComponent
}
```

**Component Properties:**
- `has_mouth` (flag): Whether entity has functional mouth
- `taste_sensitivity` (property): Taste sensitivity
- `mouth_description` (field): Description of mouth

**Default Values:**
- `has_mouth`: false
- `taste_sensitivity`: 1.0
- `mouth_description`: "expressive mouth"

#### HasBrain

Enables Consider verb and with eyes enables Close Examine.

```go
type HasBrain struct {
    component.OComponent
}
```

**Component Properties:**
- `has_brain` (flag): Whether entity has functional brain
- `intelligence` (property): Intelligence level
- `brain_type` (field): Type/description of brain

**Default Values:**
- `has_brain`: false
- `intelligence`: 10.0
- `brain_type`: "thinking brain"

#### HasSoul

Provides spiritual essence and connection.

```go
type HasSoul struct {
    component.OComponent
}
```

**Component Properties:**
- `has_soul` (flag): Whether entity has a soul
- `soul_strength` (property): Strength of spiritual essence
- `soul_type` (field): Type of soul

**Default Values:**
- `has_soul`: false
- `soul_strength`: 1.0
- `soul_type`: "mortal soul"

#### HasAttributes

Gives an object key value pairs that contain attributes meant to be rolled against.

```go
type HasAttributes struct {
    component.OComponent
}
```

**Component Properties:**
- `has_attributes` (flag): Whether entity has attributes
- `strength` (property): Physical strength
- `dexterity` (property): Agility and coordination
- `constitution` (property): Health and endurance
- `wisdom` (property): Wisdom and perception
- `charisma` (property): Social presence

**Default Values:**
- `has_attributes`: false
- `strength`: 10.0
- `dexterity`: 10.0
- `constitution`: 10.0
- `wisdom`: 10.0
- `charisma`: 10.0

#### CanRead

Enables Read verb and contains book text in the form of topics and text blocks.

```go
type CanRead struct {
    component.OComponent
}
```

**Component Properties:**
- `can_read` (flag): Whether the object can be read
- `content` (field): Main text content
- `topics` (field): Available topics/chapters
- `language` (field): Language of the text

**Default Values:**
- `can_read`: false
- `content`: "" (empty)
- `topics`: "" (empty)
- `language`: "common"

#### CanTalk

Allows an object to respond when things are said to it.

```go
type CanTalk struct {
    component.OComponent
}
```

**Component Properties:**
- `can_talk` (flag): Whether the object can engage in conversation
- `dialogue_tree` (field): Conversation structure
- `greeting` (field): Default greeting message
- `friendliness` (property): Social disposition

**Default Values:**
- `can_talk`: false
- `dialogue_tree`: "" (empty)
- `greeting`: "Hello there!"
- `friendliness`: 0.5

#### IsContainer

Enables placing objects inside of it.

```go
type IsContainer struct {
    component.OComponent
}
```

**Component Properties:**
- `is_container` (flag): Whether the object can contain other objects
- `capacity` (property): Maximum storage capacity
- `current_load` (property): Current amount stored
- `contents` (field): List of contained objects

**Default Values:**
- `is_container`: false
- `capacity`: 10.0
- `current_load`: 0.0
- `contents`: "" (empty)

#### HasPhysicalProperties

Defines Height, Width, Composition.

```go
type HasPhysicalProperties struct {
    component.OComponent
}
```

**Component Properties:**
- `has_physical_properties` (flag): Whether the object has defined physical properties
- `height` (property): Object height
- `width` (property): Object width
- `weight` (property): Object weight
- `composition` (field): Material composition

**Default Values:**
- `has_physical_properties`: false
- `height`: 1.0
- `width`: 1.0
- `weight`: 1.0
- `composition`: "unknown material"

#### HasExits

Allows an object to have multiple exits. Exits can be hidden based on object conditions.

```go
type HasExits struct {
    component.OComponent
}
```

**Component Properties:**
- `has_exits` (flag): Whether the object has exits
- `north` (field): North exit destination
- `south` (field): South exit destination
- `east` (field): East exit destination
- `west` (field): West exit destination
- `up` (field): Up exit destination
- `down` (field): Down exit destination
- `hidden_exits` (field): Hidden exit information

**Default Values:**
- `has_exits`: false
- All directional fields: "" (empty)
- `hidden_exits`: "" (empty)

#### CanHold

Enables Hold verb.

```go
type CanHold struct {
    component.OComponent
}
```

**Component Properties:**
- `can_hold` (flag): Whether the entity can hold objects
- `held_in_left` (field): Object held in left hand
- `held_in_right` (field): Object held in right hand
- `grip_strength` (property): Holding strength

**Default Values:**
- `can_hold`: false
- `held_in_left`: "" (empty)
- `held_in_right`: "" (empty)
- `grip_strength`: 1.0

#### CanWear

Enables Wear verb.

```go
type CanWear struct {
    component.OComponent
}
```

**Component Properties:**
- `can_wear` (flag): Whether the object can be worn
- `wear_location` (field): Where the object is worn
- `armor_class` (property): Defensive value
- `material` (field): Material type

**Default Values:**
- `can_wear`: false
- `wear_location`: "body"
- `armor_class`: 0.0
- `material`: "cloth"

#### CanEat

Enables Eat verb.

```go
type CanEat struct {
    component.OComponent
}
```

**Component Properties:**
- `can_eat` (flag): Whether the object can be eaten
- `nutrition_value` (property): Nutritional value
- `taste` (field): Flavor description
- `effect` (field): Effects when consumed

**Default Values:**
- `can_eat`: false
- `nutrition_value`: 1.0
- `taste`: "bland"
- `effect`: "" (empty)

#### CanDrink

Enables Drink verb.

```go
type CanDrink struct {
    component.OComponent
}
```

**Component Properties:**
- `can_drink` (flag): Whether the object can be drunk
- `liquid_amount` (property): Amount of liquid
- `liquid_type` (field): Type of liquid
- `taste` (field): Flavor description
- `effect` (field): Effects when consumed

**Default Values:**
- `can_drink`: false
- `liquid_amount`: 1.0
- `liquid_type`: "water"
- `taste`: "refreshing"
- `effect`: "" (empty)

#### CanFight

Enables combat capabilities.

```go
type CanFight struct {
    component.OComponent
}
```

**Component Properties:**
- `can_fight` (flag): Whether the entity can engage in combat
- `health` (property): Current health points
- `attack_power` (property): Attack strength
- `defense` (property): Defensive capability
- `combat_style` (field): Fighting style

**Default Values:**
- `can_fight`: false
- `health`: 100.0
- `attack_power`: 10.0
- `defense`: 5.0
- `combat_style`: "defensive"

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

#### Built-in Verbs

##### Look
Retrieves visual scope based on parameters. User can define scope by specifying what to look at. Multiple objects can be specified. Scope is sent to renderer alongside the looking object to render a top level view.

**Required Components:** HasEyes

##### Go
Attempts to move the character in the direction specified. The going object must have the necessary requirements to travel in the given direction (legs, wings, attributes).

**Required Components:** HasLegs

##### Take
Attempts to hold or add the targeted object to the taking objects container.

**Required Components:** HasHands

##### Hold
Attempts to hold an object in hand. Held items are used with actions. When fighting your held items attributes are used in calculations.

**Required Components:** HasHands, CanHold

##### Move
Moves an item from one place to another.

**Required Components:** HasHands

##### Open/Close
Opens or closes an object if it can be opened.

**Required Components:** HasHands

##### Lock/Unlock
Prevents/allows an object from being opened without a key.

**Required Components:** HasHands

##### Drop
Places an item on the ground (move overlap).

**Required Components:** HasHands

##### Listen
Takes the sounds in an active scope and sends them to the renderer.

**Required Components:** HasEars

##### Say
Sends messages into the chat stream for the most local scope in the form of heard events.

**Required Components:** HasMouth

##### Eat/Drink/Taste
Attempts to consume object, render taste and effect.

**Required Components:** HasMouth

##### Read
Attempts to read an object. Can be given a 'topics' dict that lists keys as headings and values as body text for a book, sign, etc.

**Required Components:** HasEyes

##### Wear
Attempts to wear an object.

**Required Components:** HasHands

##### Pickpocket
Attempts to covertly take from an object with a container.

**Required Components:** HasHands

##### Fight
Attempts to start a fight encounter with object.

**Required Components:** CanFight

##### Consider
Analyzes an object or situation.

**Required Components:** HasBrain

##### Close Examine
Performs detailed examination with eyes and brain.

**Required Components:** HasEyes, HasBrain

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