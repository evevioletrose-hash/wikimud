---
name: 🔴 Core Test Framework Implementation
about: Fix build issues and implement comprehensive test framework
title: "Implement Core Test Framework and Fix Build Issues"
labels: ["critical", "testing", "infrastructure", "bug"]
assignees: []
---

## 🎯 Problem
The project currently has no test files (`[no test files]` for all packages) and has a build failure in the event system:
```
internal/world/event/event.go:47:20: not enough arguments in call to obj.HandleEvent
```

## 🏆 Goals
1. Fix the immediate build/test failure in event system
2. Set up comprehensive test framework
3. Add basic test coverage for core components
4. Establish testing conventions and patterns

## ✅ Acceptance Criteria
- [ ] Fix `HandleEvent` method signature mismatch
- [ ] Create test files for each major package
- [ ] Add unit tests for existing components (HasEyes, HasLegs, CanOpen) 
- [ ] Set up test utilities and helpers
- [ ] Document testing guidelines in contributing.md
- [ ] All tests pass with `go test ./...`

## 💡 Implementation Notes
- Follow the testing examples already documented in `docs/guides/contributing.md`
- Use table-driven tests for component testing
- Focus on core functionality first: entity creation, component attachment, basic operations

## 🔗 Dependencies
None - this is foundational work that other issues depend on.

## 📝 Definition of Done
- [ ] Build passes without errors
- [ ] Test coverage for all existing components
- [ ] Testing patterns documented
- [ ] CI can run tests successfully