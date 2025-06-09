# Shared Dependencies in Go Modules

When multiple Go modules share the same dependencies, there are several important considerations and best practices to follow.

## Version Management

### Common Scenarios
- Multiple modules in a monorepo
- Shared libraries across projects
- Microservices using common packages

### Version Resolution
Go's module system automatically resolves dependency versions:
- Uses the highest compatible version
- Respects version constraints in go.mod
- Maintains consistency across modules

### Version Selection Algorithm
Go uses a specific algorithm to select dependency versions:

1. **Direct Dependencies**
   - If a module directly requires a package, it uses the version specified in its go.mod
   - Example: `require github.com/pkg/errors v0.9.1`

2. **Indirect Dependencies**
   - For dependencies of dependencies, Go selects the highest compatible version
   - Compatibility is determined by semantic versioning rules
   - A version is compatible if it's in the same major version

3. **Version Selection Rules**
   - If multiple versions are available, Go picks the highest version that satisfies all constraints
   - For example, if module A requires `v1.2.0` and module B requires `v1.3.0`, Go will use `v1.3.0`
   - If module A requires `v1.2.0` and module B requires `v2.0.0`, Go will use both versions as they're different major versions

4. **Minimal Version Selection (MVS)**
   - Go uses MVS to ensure reproducible builds
   - It selects the minimum set of versions that satisfy all requirements
   - This prevents unnecessary updates and maintains stability

5. **Version Conflicts**
   - If two modules require incompatible versions (different major versions), Go will use both
   - Each module will use its own version of the dependency
   - This is known as "multiple major version support"

Example of version selection:
```go
// Module A
require github.com/pkg/errors v1.2.0

// Module B
require github.com/pkg/errors v1.3.0

// Result: Go will use v1.3.0 as it's the highest compatible version
```

## Best Practices

### 1. Version Pinning
```go
require (
    github.com/common/package v1.2.3
)
```

### 2. Replace Directives
Use `replace` in go.mod to point to local versions:
```go
replace github.com/common/package => ./internal/common/package
```

### 3. Workspace Mode
For Go 1.18+ projects, use workspace mode:
```bash
go work init
go work use ./module1 ./module2
```

## Common Issues

### Version Conflicts
- Different modules requiring different versions
- Incompatible major versions
- Circular dependencies

### Solutions
1. Use workspace mode
2. Implement interface-based design
3. Use replace directives
4. Maintain consistent versions

## Tools

```bash
# Check for version conflicts
go mod why -m all

# Analyze dependency graph
go mod graph

# Verify workspace setup
go work sync
```

## Example Structure

```
workspace/
├── go.work
├── module1/
│   └── go.mod
├── module2/
│   └── go.mod
└── shared/
    └── go.mod
``` 