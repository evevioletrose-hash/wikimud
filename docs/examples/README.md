# WikiMUD Engine Examples

This directory contains practical examples showing how to use WikiMUD Engine to build MUD games.

## Example Index

### Basic Usage

- [Creating Entities](basic-entity-creation.md) - How to create and configure game objects
- [Component Usage](component-examples.md) - Working with the component system  
- [Prototype Definitions](prototype-examples.md) - JSON-based object templates
- [Scope Management](scope-examples.md) - Organizing entities in rooms/areas

### Advanced Features

- [Custom Components](custom-component.md) - Building your own components
- [Verb Implementation](custom-verb.md) - Adding new player commands
- [Rendering Systems](renderer-examples.md) - Different output formats
- [Event Handling](event-examples.md) - Using the event system

### Complete Examples

- [Simple Room](simple-room.md) - A basic MUD room with items
- [Interactive NPC](interactive-npc.md) - Non-player character with behaviors
- [Container System](container-system.md) - Bags, chests, and item storage
- [Combat System](combat-system.md) - Basic fighting mechanics

## Running Examples

Most examples are self-contained Go files that can be run directly:

```bash
cd docs/examples
go run basic-entity-creation.go
```

Some examples require copying code into your own `main.go` file or integrating with the existing project structure.

## Example Structure

Each example follows this pattern:

1. **Overview**: What the example demonstrates
2. **Code**: Complete, runnable code
3. **Explanation**: Step-by-step breakdown
4. **Extensions**: Ideas for building on the example
5. **Related**: Links to related examples and documentation

## Getting Started

If you're new to WikiMUD Engine, start with:

1. [Creating Entities](basic-entity-creation.md)
2. [Component Usage](component-examples.md)  
3. [Simple Room](simple-room.md)
4. [Custom Components](custom-component.md)

Then explore the other examples based on your specific needs.