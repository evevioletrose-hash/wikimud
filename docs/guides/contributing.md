# Contributing to WikiMUD Engine

Thank you for your interest in contributing to WikiMUD Engine! This guide will help you get started with contributing to the project.

## Getting Started

1. **Fork the repository** on GitHub
2. **Clone your fork** locally:
   ```bash
   git clone https://github.com/your-username/wikimud.git
   cd wikimud
   ```
3. **Set up the development environment** following the [Developer Guide](developer.md)
4. **Create a feature branch**:
   ```bash
   git checkout -b feature/your-feature-name
   ```

## Ways to Contribute

### 🐛 Bug Reports

Help us improve by reporting bugs:

- **Search existing issues** first to avoid duplicates
- **Use the bug report template** with clear steps to reproduce
- **Include system information** (Go version, OS, etc.)
- **Provide minimal test cases** when possible

### 💡 Feature Requests

Suggest new features or improvements:

- **Check the roadmap** and existing feature requests
- **Describe the use case** and motivation
- **Consider implementation complexity** and maintenance
- **Discuss breaking changes** with maintainers first

### 📝 Documentation

Improve documentation for users and developers:

- **Fix typos and unclear explanations**
- **Add missing API documentation**
- **Create tutorials and examples**
- **Improve inline code comments**

### 🔧 Code Contributions

#### Types of Code Contributions

1. **New Components**: Extend entity capabilities
2. **New Verbs**: Add player actions and commands  
3. **New Renderers**: Support different output formats
4. **Bug Fixes**: Fix existing issues
5. **Performance Improvements**: Optimize existing code
6. **Refactoring**: Improve code structure and maintainability

#### Before You Start

- **Discuss major changes** in issues before implementing
- **Check existing code** for similar functionality
- **Follow existing patterns** and conventions
- **Write tests** for new functionality

## Development Guidelines

### Code Style

We follow standard Go conventions:

```bash
# Format code
go fmt ./...

# Check for common issues  
go vet ./...

# Run static analysis (if available)
golint ./...
```

#### Naming Conventions

- **Packages**: lowercase, single word when possible
- **Types**: PascalCase (e.g., `WorldObject`, `ComponentManager`)
- **Functions/Methods**: PascalCase for exported, camelCase for internal
- **Variables**: camelCase
- **Constants**: PascalCase or UPPER_CASE for package-level

#### Documentation

All exported functions and types must have Go doc comments:

```go
// NewWorldObject creates a new entity with the specified properties.
// The prototype parameter defines the initial component configuration.
func NewWorldObject(id string, prototype *ObjectPrototype) *WorldObject {
    // implementation
}
```

### Component Development

#### Component Checklist

When creating a new component:

- [ ] **Embed `OComponent`** as base structure
- [ ] **Implement `BuildComponent`** method  
- [ ] **Provide default values** for all properties
- [ ] **Handle missing prototype data** gracefully
- [ ] **Document component properties** in comments
- [ ] **Add unit tests** for component behavior
- [ ] **Update documentation** with new component

#### Component Example

```go
package components

import (
    "mud-engine/internal/world/entity/component"
    "mud-engine/internal/world/entity/prototype"
)

// HasWings provides flight capabilities to entities.
// 
// Properties:
//   - can_fly (flag): Whether the entity can fly
//   - wing_span (property): Wing span in meters
//   - wing_description (field): Description of wings
type HasWings struct {
    component.OComponent
}

// BuildComponent initializes a HasWings component from prototype data.
func (h *HasWings) BuildComponent(p *prototype.ObjectPrototype) component.OComponent {
    newComponent := HasWings{}
    oComponent := newComponent.OComponent
    
    // Initialize flight capability
    if canFly, ok := p.Flags["can_fly"]; ok {
        oComponent.SetFlag("can_fly", canFly)
    } else {
        oComponent.SetFlag("can_fly", false)
    }
    
    // Initialize wing span
    if wingSpan, ok := p.Properties["wing_span"]; ok {
        oComponent.SetProperty("wing_span", wingSpan)
    } else {
        oComponent.SetProperty("wing_span", 1.0)
    }
    
    // Initialize wing description
    if desc, ok := p.Fields["wing_description"]; ok {
        oComponent.SetField("wing_description", desc)
    } else {
        oComponent.SetField("wing_description", "functional wings")
    }
    
    return newComponent.OComponent
}
```

### Testing

#### Test Structure

```
component_test.go      # Component unit tests
integration_test.go    # Cross-component tests
examples_test.go       # Example usage tests
```

#### Test Examples

```go
func TestHasWingsComponent(t *testing.T) {
    tests := []struct {
        name      string
        prototype *prototype.ObjectPrototype
        expected  map[string]interface{}
    }{
        {
            name: "default values",
            prototype: &prototype.ObjectPrototype{
                Fields:     make(map[string]string),
                Flags:      make(map[string]bool),
                Properties: make(map[string]float32),
            },
            expected: map[string]interface{}{
                "can_fly":          false,
                "wing_span":        float32(1.0),
                "wing_description": "functional wings",
            },
        },
        {
            name: "custom values",
            prototype: &prototype.ObjectPrototype{
                Fields:     map[string]string{"wing_description": "majestic wings"},
                Flags:      map[string]bool{"can_fly": true},
                Properties: map[string]float32{"wing_span": 3.5},
            },
            expected: map[string]interface{}{
                "can_fly":          true,
                "wing_span":        float32(3.5),
                "wing_description": "majestic wings",
            },
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            hasWings := &HasWings{}
            component := hasWings.BuildComponent(tt.prototype)
            
            // Test each expected value
            if expected, ok := tt.expected["can_fly"].(bool); ok {
                actual, exists := component.GetFlag("can_fly")
                assert.True(t, exists, "can_fly flag should exist")
                assert.Equal(t, expected, actual)
            }
            
            // ... test other properties
        })
    }
}
```

### Commit Guidelines

#### Commit Message Format

```
type(scope): short description

Longer description if needed explaining what and why.

Fixes #123
```

#### Types

- **feat**: New feature
- **fix**: Bug fix
- **docs**: Documentation changes
- **style**: Code style changes (formatting, etc.)
- **refactor**: Code refactoring without feature changes
- **test**: Adding or updating tests
- **chore**: Maintenance tasks

#### Scopes

- **component**: Component-related changes
- **verb**: Verb system changes
- **render**: Rendering system changes
- **entity**: Entity system changes
- **scope**: Scope management changes
- **event**: Event system changes
- **docs**: Documentation changes

#### Examples

```
feat(component): add HasWings component for flying entities

Adds a new component that enables entities to have flight capabilities.
Includes wing span property and flight status flag.

Fixes #45

---

fix(verb): resolve target validation in look command

The look verb was incorrectly validating targets when multiple objects
had similar names. Updated to use exact match first, then fuzzy matching.

Fixes #67

---

docs(api): add comprehensive component documentation

- Added component property tables
- Included usage examples
- Documented default values and behavior
```

## Pull Request Process

### Before Submitting

1. **Run tests**: `go test ./...`
2. **Format code**: `go fmt ./...`
3. **Check for issues**: `go vet ./...`
4. **Update documentation** if needed
5. **Add tests** for new functionality

### Pull Request Template

```markdown
## Description
Brief description of changes.

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Documentation update
- [ ] Refactoring
- [ ] Performance improvement

## Testing
- [ ] Unit tests added/updated
- [ ] Integration tests pass
- [ ] Manual testing completed

## Documentation
- [ ] Code comments updated
- [ ] API documentation updated
- [ ] User documentation updated

## Breaking Changes
List any breaking changes and migration notes.

## Related Issues
Closes #123
References #456
```

### Review Process

1. **Automated checks** must pass (builds, tests)
2. **Code review** by maintainer(s)
3. **Address feedback** and update as needed
4. **Final approval** and merge

## Code Review Checklist

### For Reviewers

- [ ] **Code follows project conventions**
- [ ] **Tests are comprehensive and pass**
- [ ] **Documentation is updated**
- [ ] **No breaking changes without discussion**
- [ ] **Performance impact is acceptable**
- [ ] **Security considerations addressed**

### For Contributors

- [ ] **Self-review completed**
- [ ] **Edge cases considered**
- [ ] **Error handling implemented**
- [ ] **Memory leaks avoided**
- [ ] **Concurrency safety considered**

## Community Guidelines

### Be Respectful

- **Use inclusive language**
- **Be constructive in feedback**
- **Respect different perspectives**
- **Help newcomers learn**

### Communication

- **Be clear and specific** in issue reports
- **Provide context** for feature requests
- **Ask questions** when unsure
- **Share knowledge** with the community

### Recognition

Contributors are recognized in:

- **CONTRIBUTORS.md** file
- **Release notes** for significant contributions
- **GitHub contribution graph**
- **Special thanks** in documentation

## Release Process

### Version Numbering

We follow [Semantic Versioning](https://semver.org/):

- **MAJOR**: Breaking changes
- **MINOR**: New features, backward compatible
- **PATCH**: Bug fixes, backward compatible

### Release Checklist

- [ ] **Update version numbers**
- [ ] **Update CHANGELOG.md**
- [ ] **Run full test suite**
- [ ] **Build and test release artifacts**
- [ ] **Update documentation**
- [ ] **Create release notes**
- [ ] **Tag release in Git**

## Getting Help

### Resources

- **Documentation**: Check `docs/` directory
- **Examples**: Review existing code patterns
- **Issues**: Search existing GitHub issues
- **Discussions**: Use GitHub Discussions for questions

### Contact

- **GitHub Issues**: Bug reports and feature requests
- **GitHub Discussions**: General questions and ideas
- **Pull Requests**: Code contributions and reviews

Thank you for contributing to WikiMUD Engine! 🎮