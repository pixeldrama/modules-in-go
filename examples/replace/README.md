# Go Module Replace Example

This example demonstrates how to use the `replace` directive in Go modules to replace a module dependency with a local version.

## Structure

- `module1/`: A simple module that provides a Hello function
- `module2/`: A module that depends on module1 but uses replace to point to the local version

## How to Run

1. Navigate to the module2 directory:
   ```bash
   cd module2
   ```

2. Run the example:
   ```bash
   go run main.go
   ```

You should see the output: "Hello from module1"

## How it Works

The `replace` directive in `module2/go.mod` tells Go to use the local version of module1 instead of trying to download it from GitHub:

```
replace github.com/example/module1 => ../module1
```

This is useful for:
- Local development
- Testing changes to dependencies
- Working with forks
- Avoiding publishing modules to a repository during development 