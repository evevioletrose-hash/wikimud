# Developer Setup Guide

This guide will help you set up a development environment for WikiMUD Engine.

## Prerequisites

- **Go 1.20+**: Download from [golang.org](https://golang.org/dl/)
- **Git**: For version control
- **Code Editor**: VS Code, GoLand, or any Go-compatible editor

## Setup

### 1. Clone the Repository

```bash
git clone https://github.com/evevioletrose-hash/wikimud.git
cd wikimud
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Build the Project

```bash
go build
```

### 4. Run Tests (when available)

```bash
go test ./...
```

### 5. Run the Example

```bash
go run main.go
```

## Development Workflow

### Project Structure

```
wikimud/
├── main.go                 # Example application
├── go.mod                  # Go module definition
├── internal/               # Private application code
│   ├── app/               # Application logic
│   └── world/             # MUD engine core
│       ├── entity/        # Entity-Component system
│       │   ├── entity.go
│       │   ├── component/
│       │   ├── prototype/
│       │   ├── verb/
│       │   └── attribute/
│       ├── scope/         # Entity scoping
│       ├── render/        # Output rendering
│       ├── event/         # Event system
│       └── encounter/     # Game encounters
└── docs/                  # Documentation
```

### Code Organization

#### Internal Packages

- **`internal/app/`**: Main application entry point and coordination
- **`internal/world/`**: Core MUD engine implementation
- **`internal/world/entity/`**: Entity-Component-System implementation
- **`internal/world/scope/`**: Entity scope management for rooms/areas
- **`internal/world/render/`**: Text rendering and formatting
- **`internal/world/event/`**: Event handling and routing

#### Import Paths

All internal packages use the module name `mud-engine`:

```go
import (
    "mud-engine/internal/world/entity"
    "mud-engine/internal/world/entity/component"
    "mud-engine/internal/world/scope"
)
```

### Adding New Components

1. **Create Component File**:

```bash
touch internal/world/entity/component/components/my-component.go
```

2. **Implement Component Structure**:

```go
package components

import (
    "mud-engine/internal/world/entity/component"
    "mud-engine/internal/world/entity/prototype"
)

type MyComponent struct {
    component.OComponent
}

func (c *MyComponent) BuildComponent(p *prototype.ObjectPrototype) component.OComponent {
    newComponent := MyComponent{}
    oComponent := newComponent.OComponent
    
    // Initialize component fields, flags, and properties
    if val, ok := p.Fields["my_field"]; ok {
        oComponent.SetField("my_field", val)
    } else {
        oComponent.SetField("my_field", "default_value")
    }
    
    return newComponent.OComponent
}
```

3. **Register Component** (implementation varies by registration system)

### Adding New Verbs

1. **Define Verb Logic** in existing verb system:

```go
type MyVerb struct {
    Name string
}

func (v *MyVerb) Execute(actor entity.WorldObject, targets []entity.WorldObject, params []string) string {
    // Implement verb logic
    return "Action completed"
}
```

2. **Register with World**:

```go
world.register_verb(verb.VerbType{Name: "myverb"})
```

### Adding New Renderers

1. **Create Renderer** in `internal/world/render/renderers/`:

```go
package renderers

import "mud-engine/internal/world/scope"

type MyRenderer struct{}

func (r *MyRenderer) Render(scope scope.Scope) string {
    // Implement custom rendering logic
    return "Rendered output"
}
```

2. **Use Renderer**:

```go
renderer := &renderers.MyRenderer{}
output := renderer.Render(gameScope)
```

## Development Best Practices

### Code Style

- Follow Go conventions (gofmt, golint)
- Use descriptive variable names
- Add comments for exported functions
- Keep functions small and focused
- Use interfaces for abstraction

### Component Design

- **Single Responsibility**: Each component should handle one aspect
- **Composition over Inheritance**: Use component composition
- **Default Values**: Always provide sensible defaults
- **Type Safety**: Use appropriate Go types (string, bool, float32)

### Entity Management

- **Unique IDs**: Ensure entity IDs are unique within their scope
- **Lifecycle Management**: Properly initialize and cleanup entities
- **Component Dependencies**: Document component interactions
- **Event Handling**: Use events for loose coupling

### Testing Strategy

- **Unit Tests**: Test individual components and functions
- **Integration Tests**: Test component interactions
- **Game Logic Tests**: Test complete verb execution flows
- **Prototype Tests**: Test JSON prototype parsing

Example test structure:

```go
func TestHasEyesComponent(t *testing.T) {
    prototype := &prototype.ObjectPrototype{
        Fields: map[string]string{"eye_color": "blue"},
        Properties: map[string]float32{"eye_count": 2},
        Flags: map[string]bool{"has_eyes": true},
    }
    
    hasEyes := &components.HasEyes{}
    component := hasEyes.BuildComponent(prototype)
    
    eyeColor, exists := component.GetField("eye_color")
    assert.True(t, exists)
    assert.Equal(t, "blue", eyeColor)
}
```

### Debugging

#### Common Issues

1. **Missing Component Properties**: Check prototype definitions
2. **Entity Not Found**: Verify entity is added to correct scope
3. **Verb Execution Failures**: Check target validation logic
4. **Import Errors**: Ensure correct module paths

#### Debugging Tools

```go
// Debug entity state
fmt.Printf("Entity: %+v\n", entity)
fmt.Printf("Components: %+v\n", entity.Components())

// Debug component properties
for _, comp := range entity.Components() {
    fmt.Printf("Component: %s\n", comp.Name)
    fmt.Printf("Fields: %+v\n", comp.Fields)
    fmt.Printf("Flags: %+v\n", comp.Flags)
    fmt.Printf("Properties: %+v\n", comp.Properties)
}
```

### Performance Optimization

#### Entity Management

- Use entity pools for frequently created/destroyed objects
- Cache component lookups where appropriate
- Minimize scope traversals
- Batch entity operations when possible

#### Component Design

- Use appropriate data types (avoid interface{} when possible)
- Implement lazy initialization for expensive components
- Consider component sharing for identical data

#### Rendering

- Cache rendered output when entity state hasn't changed
- Use string builders for complex text assembly
- Implement partial rendering for large scopes

## IDE Setup

### VS Code

Install recommended extensions:

```json
{
    "recommendations": [
        "golang.go",
        "ms-vscode.vscode-json"
    ]
}
```

### GoLand

- Enable Go modules support
- Configure code style to use gofmt
- Set up run configurations for main.go

## Git Workflow

### Branch Naming

- `feature/component-name` for new components
- `feature/verb-name` for new verbs  
- `fix/issue-description` for bug fixes
- `docs/section-name` for documentation

### Commit Messages

Follow conventional commit format:

```
type(scope): description

Examples:
feat(component): add HasWings component for flying entities
fix(verb): resolve target validation in complex verbs  
docs(api): add HasEyes component documentation
refactor(entity): simplify component initialization
```

### Pull Request Process

1. Create feature branch from main
2. Implement changes with tests
3. Update documentation
4. Submit pull request with description
5. Address review feedback
6. Merge after approval

## Troubleshooting

### Common Build Issues

**Import path errors:**
```bash
go mod tidy
go clean -modcache
```

**Missing dependencies:**
```bash
go get ./...
```

**Module path conflicts:**
Ensure all imports use `mud-engine` as module name.

### Runtime Issues

**Entity not found:**
- Verify entity was added to correct scope
- Check entity ID spelling/casing

**Component data missing:**
- Verify prototype contains required fields
- Check component BuildComponent implementation

**Verb execution fails:**
- Validate target entities exist
- Check verb parameter parsing
- Ensure actor has required permissions

## Getting Help

- Check existing issues on GitHub
- Review documentation in `docs/`
- Examine example code in repository
- Ask questions in project discussions