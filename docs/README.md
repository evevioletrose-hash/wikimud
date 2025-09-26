# WikiMUD Engine Documentation

Welcome to the comprehensive documentation for WikiMUD Engine, a sophisticated Multi-User Dungeon (MUD) game engine built on modern Entity-Component-System architecture.

## Quick Navigation

### 🚀 Getting Started
- [Main README](../README.md) - Project overview and quick start
- [Developer Setup](guides/developer.md) - Development environment setup
- [Basic Entity Creation](examples/basic-entity-creation.md) - Your first entities

### 📚 Core Documentation
- [Architecture Guide](architecture/README.md) - Understanding the ECS design
- [API Reference](api/README.md) - Complete API documentation
- [Contributing Guide](guides/contributing.md) - How to contribute

### 💡 Examples & Tutorials
- [Examples Overview](examples/README.md) - All available examples
- [Simple Room](examples/simple-room.md) - Complete MUD room
- [Custom Components](examples/custom-component.md) - Building your own components

## Documentation Structure

```
docs/
├── README.md                    # This file - documentation index
├── architecture/                # System design and architecture
│   └── README.md               # ECS architecture deep dive
├── api/                        # API reference documentation  
│   └── README.md               # Complete API documentation
├── guides/                     # Development and contribution guides
│   ├── developer.md            # Setup and development workflow
│   └── contributing.md         # Contribution guidelines  
└── examples/                   # Practical examples and tutorials
    ├── README.md               # Examples index
    ├── basic-entity-creation.md # Entity creation fundamentals
    ├── simple-room.md          # Complete room example
    └── custom-component.md     # Building custom components
```

## Key Concepts

### Entity-Component-System (ECS)
WikiMUD uses ECS architecture for maximum flexibility:
- **Entities**: Unique game objects (players, items, NPCs)
- **Components**: Modular data containers (HasEyes, CanOpen, etc.)
- **Systems**: Logic that operates on entities with specific components

### Core Systems

#### 🎯 **Entity Management**
- `WorldObject` - The fundamental entity type
- `ComponentManager` - Orchestrates component lifecycle
- `Scope` - Containers for organizing entities (rooms, areas)

#### 🧩 **Component System**
- `OComponent` - Base component with flags, fields, and properties
- Built-in components: `HasEyes`, `HasLegs`, `CanOpen`
- Custom component support with `BuildComponent` pattern

#### 🎨 **Rendering**
- Multiple output formats (raw text, linked text, journal view)
- `Renderer` interface for extensible output systems
- Scope-to-text conversion for player interfaces

#### ⚡ **Actions & Events**
- `Verb` system for player commands and actions
- Event emission and handling for reactive programming
- Flexible parameter passing and target validation

#### 🏗️ **Prototypes**
- JSON-based entity templates
- Automatic type inference (string → field, boolean → flag, number → property)
- Reusable object definitions

## Learning Path

### Beginner (New to WikiMUD)
1. Read [Project Overview](../README.md)
2. Follow [Developer Setup](guides/developer.md)
3. Try [Basic Entity Creation](examples/basic-entity-creation.md)
4. Build a [Simple Room](examples/simple-room.md)

### Intermediate (Ready to Extend)
1. Study [Architecture Guide](architecture/README.md)
2. Create [Custom Components](examples/custom-component.md)
3. Review [API Reference](api/README.md)
4. Explore built-in component implementations

### Advanced (Contributing to Engine)
1. Read [Contributing Guidelines](guides/contributing.md)
2. Study existing codebase patterns
3. Implement new systems (renderers, verbs, components)
4. Add tests and documentation

## Feature Highlights

### 🎮 **Game Development Ready**
- Complete entity lifecycle management
- Flexible component composition
- Multi-format rendering support
- Event-driven architecture

### 🔧 **Developer Friendly**
- Clean Go API with full documentation
- JSON-based configuration
- Modular, testable components
- Extensive examples and tutorials

### 🚀 **Performance Focused**
- Efficient entity lookup (O(1) by ID)
- Component-based data organization
- Minimal memory allocation patterns
- Scope isolation for scalability

### 📈 **Extensible Design**
- Plugin-style component system
- Custom renderer support
- Flexible verb implementation
- Event system for loose coupling

## Common Use Cases

### 🏰 **Traditional MUD**
- Text-based multiplayer adventure
- Room-based world exploration
- Character progression systems
- Item and inventory management

### 📖 **Interactive Fiction**
- Story-driven experiences
- Rich narrative descriptions
- Character-driven interactions
- Branching storylines

### 🎲 **Game Prototyping**
- Rapid gameplay iteration
- Component-based feature testing
- Flexible rule systems
- Easy content creation

### 🎯 **Educational Projects**
- Game architecture learning
- Entity-Component-System demonstrations
- Go programming examples
- Software design patterns

## API Quick Reference

### Entity Creation
```go
prototype, _ := prototype.NewObjectPrototypeFromJSON(jsonStr)
entity := entity.NewEntity("unique_id", prototype)
```

### Scope Management
```go
room := scope.NewScope()
room.AddEntity("entity_id", entity)
entity, exists := room.GetEntity("entity_id")
```

### Component Access
```go
components := entity.Components()
for _, comp := range components {
    value, exists := comp.GetProperty("property_name")
    field, exists := comp.GetField("field_name")
    flag, exists := comp.GetFlag("flag_name")
}
```

### Verb Execution
```go
verb := verb.VerbType{Name: "look"}
action := &verb.Verb{}
result := action.Try(actor, verb, targets, params)
```

## Support & Community

### 📖 **Documentation**
- Comprehensive guides for all skill levels
- API reference with examples
- Architecture explanations
- Best practices and patterns

### 🛠️ **Development**
- Clean, well-documented Go code
- Extensive test coverage (when available)
- Consistent code style and conventions
- Regular updates and improvements

### 🤝 **Contributing**
- Welcoming to new contributors
- Clear contribution guidelines
- Code review process
- Recognition for contributors

## Next Steps

Choose your path based on your goals:

- **Want to build a MUD?** → Start with [Basic Entity Creation](examples/basic-entity-creation.md)
- **Want to understand the system?** → Read [Architecture Guide](architecture/README.md)
- **Want to extend the engine?** → Study [Custom Components](examples/custom-component.md)
- **Want to contribute?** → Follow [Contributing Guide](guides/contributing.md)

---

**WikiMUD Engine** - Powerful, flexible, and developer-friendly MUD engine for the modern era.