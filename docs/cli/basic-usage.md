# Basic Usage of Cobra

This guide shows how to create a simple command-line application using the Cobra library.

## Setup

Create a new project and install Cobra:

```bash
mkdir hello-cli
cd hello-cli
go mod init hello-cli
go get -u github.com/spf13/cobra@latest
```

## Simple Command

Create a basic `main.go`:

```go
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	// Create the root command
	rootCmd := &cobra.Command{
		Use:   "hello", // The name of the binary
		Short: "A simple CLI application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, World!")
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```

### Understanding the Code

- **cobra.Command**: This is the core structure in Cobra. It defines a command with its behavior.
  - `Use`: Defines the command name as it appears in help text
  - `Short`: A brief description shown in help text
  - `Run`: A function that executes when the command runs

- **rootCmd.Execute()**: This method parses command-line arguments and executes the appropriate command. It's always called on the root command, even when you have subcommands.

Build and run:

```bash
go build -o hello
./hello
```

Output:
```
Hello, World!
```

## Using Flags

Add flags to your root command:

```go
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var name string

	// Create the root command
	rootCmd := &cobra.Command{
		Use:   "hello", // The name of the binary
		Short: "A simple CLI application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Hello, %s!\n", name)
		},
	}

	// Add flags to the root command
	rootCmd.Flags().StringVarP(&name, "name", "n", "World", "Name to greet")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```

### Understanding Flags

Flags are command-line options that modify a command's behavior:

- **rootCmd.Flags()**: Accesses the flag set for this command
- **StringVarP()**: Defines a string flag with both long and short forms
  - Parameters:
    - `&name`: Variable to store the flag value
    - `"name"`: Long form of the flag (used as `--name`)
    - `"n"`: Short form of the flag (used as `-n`)
    - `"World"`: Default value if flag is not provided
    - `"Name to greet"`: Help text describing the flag

Cobra also provides other flag types like `BoolP()`, `IntP()`, and `Float64P()`.

Example usage:

```bash
$ ./hello --name Alice
Hello, Alice!

$ ./hello -n Bob
Hello, Bob!

$ ./hello
Hello, World!
```

## Using Arguments

Create a command that uses arguments:

```go
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	// Create the root command
	rootCmd := &cobra.Command{
		Use:   "hello [text to echo]", // The binary name followed by argument description
		Short: "A simple CLI application",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(strings.Join(args, " "))
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```

### Understanding Arguments

Arguments are values passed to a command without flags:

- **Use**: The `[text to echo]` part describes the expected arguments
- **Args**: Validates the number of arguments provided
  - `cobra.MinimumNArgs(1)`: Requires at least one argument
  - Other validators: `ExactArgs(n)`, `MaximumNArgs(n)`, `RangeArgs(min, max)`
- **args parameter**: In the `Run` function, `args` is a slice containing all arguments passed to the command

Example usage:

```bash
$ ./hello Hello World
Hello World
```

## Understanding the "Use" Field

The `Use` field in Cobra commands defines how the command appears in help text and how it's used in the command line:

```go
rootCmd := &cobra.Command{
    Use:   "hello", // The binary name
    // ...
}

rootCmd := &cobra.Command{
    Use:   "hello [text to echo]", // The binary name followed by argument description
    // ...
}
```

- For a **root command**, `Use` should be the binary name (e.g., "hello")
- You can optionally add argument descriptions after the binary name (e.g., "hello [text to echo]")

## Automatic Help

Cobra generates help text automatically:

```bash
$ ./hello --help
A simple CLI application

Usage:
  hello [flags]

Flags:
  -h, --help          help for hello
  -n, --name string   Name to greet (default "World")
```

## Next Steps

For more advanced features like subcommands and persistent flags, see:

- [Advanced Usage](advanced-usage.md)
