# Go Modules

Go modules are the standard dependency management system for Go projects. They provide a way to manage dependencies, versioning, and reproducible builds.

## Key Concepts

### Module Definition
A Go module is defined by a `go.mod` file in the root directory of your project. This file:
- Declares the module path
- Lists all dependencies with their versions
- Specifies the Go version used

### Module Path
The module path is the import path prefix for all packages in the module. It typically follows the format:
```
github.com/username/project
```

### Versioning
Go modules use semantic versioning (semver):
- v0.x.x: Development versions
- v1.x.x: Stable versions
- v2.x.x: Major version updates

## Common Commands

```bash
# Initialize a new module
go mod init module-name

# Add missing and remove unused modules
go mod tidy

# Download dependencies
go mod download

# Verify dependencies
go mod verify

# List all dependencies
go list -m all
```

## Best Practices

1. Always commit `go.mod` and `go.sum` files
2. Use `go mod tidy` before committing
3. Specify exact versions for dependencies
4. Keep dependencies up to date
5. Use semantic versioning for your modules

## Module Structure

```
project/
├── go.mod
├── go.sum
├── main.go
└── internal/
    └── package/
```

## Private Modules

For private modules:
1. Configure Git to use HTTPS
2. Set GOPRIVATE environment variable
3. Configure Git credentials

## Upgrading Dependencies

```bash
# Upgrade to latest minor or patch version
go get -u ./...

# Upgrade to latest major version
go get -u -t ./...
```

## Additional Resources

- [Managing Shared Dependencies](shared-dependencies.md) - Learn about handling dependencies across multiple modules
- [Module Versioning](versioning.md) - Detailed guide on version types, constraints, and best practices 