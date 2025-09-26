# WikiMUD Engine

WikiMUD is a sophisticated Multi-User Dungeon (MUD) game engine written in Go, featuring a modern Entity-Component-System (ECS) architecture designed for building text-based multiplayer games.

## Features

- **Entity-Component-System Architecture**: Flexible, modular design for game objects
- **Prototype-Based Object Creation**: JSON-driven object definitions
- **Multi-Format Rendering**: Support for raw text, linked text, and journal views
- **Verb-Based Action System**: Extensible command system for player interactions
- **Event-Driven Architecture**: Decoupled event handling for game logic
- **Scope Management**: Hierarchical entity organization (rooms, areas, etc.)
- **Component Library**: Pre-built components for common MUD features

## Quick Start

```bash
# Clone the repository
git clone https://github.com/evevioletrose-hash/wikimud.git
cd wikimud

# Build the engine
go build

# Run the example
go run main.go
```

## Architecture Overview

```
WikiMUD Engine
├── World                    # Top-level game world container
├── Scopes                   # Contextual entity containers (rooms, areas)
├── Entities                 # Game objects (players, items, NPCs)
│   ├── Components           # Modular behavior units
│   ├── Prototypes          # Object templates
│   └── Attributes          # Entity properties
├── Verbs                   # Player actions and commands
├── Events                  # Game event system
└── Renderers              # Output formatting
```

## Project Structure

- `internal/app/`: Application entry point and main logic
- `internal/world/`: Core MUD engine components
  - `entity/`: Entity-Component-System implementation
  - `scope/`: Entity scope management
  - `render/`: Output rendering system
  - `event/`: Event handling system
- `docs/`: Comprehensive documentation
- `main.go`: Example application

## Documentation

- [Architecture Guide](docs/architecture/README.md) - Understanding the ECS design
- [API Reference](docs/api/README.md) - Complete API documentation
- [Developer Guide](docs/guides/developer.md) - Setup and development workflow
- [Examples](docs/examples/) - Code examples and tutorials

## Components

WikiMUD includes several built-in components:

- **HasEyes**: Vision and sight mechanics
- **HasLegs**: Movement and locomotion
- **CanOpen**: Container and door interactions

## Rendering Systems

- **Raw Text**: Plain text output
- **Linked Text**: Hyperlinked interactive text
- **Journal View**: Structured narrative format

## Development Status & Roadmap

WikiMUD Engine is in active development. See the [Development Roadmap](DEVELOPMENT_ROADMAP.md) for current priorities and planned features.

**Current Status:**
- ✅ Core architecture implemented
- ✅ Basic component system working  
- ❌ Test framework needed (critical priority)
- ❌ Multi-user networking not yet implemented
- ❌ Game mechanics need development

## Contributing

Contributions are welcome! Please check the [Development Roadmap](DEVELOPMENT_ROADMAP.md) for current priorities, then see our [Contributing Guide](docs/guides/contributing.md) for development setup and guidelines.

## License

This project is open source. See LICENSE file for details.
