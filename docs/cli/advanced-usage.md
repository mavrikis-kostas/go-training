# Advanced Usage of Cobra

This guide demonstrates advanced features of the Cobra library for building command-line applications.

## Introduction to Subcommands

One of the most powerful features of Cobra is the ability to create subcommands. While the basic usage focuses on a
single root command, real-world CLI applications often need multiple commands with different functionality.

For example, Git uses subcommands like `git clone`, `git commit`, and `git push`. With Cobra, you can create a similar
command structure for your application.

### How Subcommands Work

Subcommands in Cobra are simply additional `cobra.Command` instances that are attached to a parent command:

1. Create a root command (the parent)
2. Create subcommands (the children)
3. Add subcommands to the root command using `rootCmd.AddCommand(subCmd)`
4. When a user runs your application, Cobra matches the command-line arguments to the appropriate command

## Command Structure

A typical project structure for a Cobra application with multiple commands:

```
hello-cli/
├── cmd/
│   ├── root.go
│   ├── greet.go
│   ├── echo.go
│   └── version.go
└── main.go
```

### Root Command

The root command in `cmd/root.go`:

```go
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "hello", // The name of the binary
		Short: "A simple CLI application",
		Long:  "A simple CLI application built with Cobra.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Use --help to see available commands")
		},
	}

	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose output")

	return rootCmd
}

func Execute() {
	rootCmd := RootCmd()

	// Add all subcommands
	rootCmd.AddCommand(GreetCmd())
	rootCmd.AddCommand(EchoCmd())
	rootCmd.AddCommand(VersionCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```

### Flag-Based Command

A command that uses flags in `cmd/greet.go`:

```go
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func GreetCmd() *cobra.Command {
	var name string

	greetCmd := &cobra.Command{
		Use:     "greet", // The name of the subcommand
		Aliases: []string{"hello", "hi"},
		Short:   "Greet a person by name",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Hello, %s!\n", name)

			verbose, _ := cmd.Flags().GetBool("verbose")
			if verbose {
				fmt.Println("Verbose mode enabled")
			}
		},
	}

	greetCmd.Flags().StringVarP(&name, "name", "n", "World", "Name to greet")

	return greetCmd
}
```

### Argument-Based Command

A command that uses arguments in `cmd/echo.go`:

```go
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func EchoCmd() *cobra.Command {
	echoCmd := &cobra.Command{
		Use:   "echo [string to echo]", // The name of the subcommand followed by argument description
		Short: "Echo anything to the screen",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(strings.Join(args, " "))

			verbose, _ := cmd.Flags().GetBool("verbose")
			if verbose {
				fmt.Printf("Echoed %d arguments\n", len(args))
			}
		},
	}

	return echoCmd
}
```

### Version Command

A simple version command in `cmd/version.go`:

```go
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func VersionCmd() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version", // The name of the subcommand
		Short: "Print version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hello-cli v1.0.0")
		},
	}

	return versionCmd
}
```

### Main Function

The main.go file:

```go
package main

import "hello-cli/cmd"

func main() {
	cmd.Execute()
}
```

### Example Usage

```bash
# Using the flag-based command
$ ./hello greet --name Alice
Hello, Alice!

# Using the argument-based command
$ ./hello echo This is an example
This is an example

# Using command alias
$ ./hello hi --name Bob
Hello, Bob!

# With verbose flag
$ ./hello echo Hello world --verbose
Hello world
Echoed 2 arguments

# Version command
$ ./hello version
hello-cli v1.0.0
```

## Key Cobra Features

### Working with Arguments

Cobra provides several ways to handle command arguments:

```go
func EchoCmd() *cobra.Command {
return &cobra.Command{
Use:   "echo [string to echo]", // Describes expected arguments in the 'help' output
Args:  cobra.MinimumNArgs(1), // Validates argument count
Run: func (cmd *cobra.Command, args []string) {
// Access arguments directly from the args slice
fmt.Println(strings.Join(args, " "))
},
}
}
```

#### Argument Validation

Cobra provides built-in validators to ensure users provide the correct number of arguments.

For example, if you want to ensure that a command has exactly one argument, you can use `ExactArgs(1)`:

```bash
$ ./hello echo one two    # Error if ExactArgs(1)
$ ./hello echo one        # OK if ExactArgs(1)
```

Other validators include:

- `MinimumNArgs(n)`: Requires at least `n` arguments
- `MaximumNArgs(n)`: Allows up to `n` arguments
- `RangeArgs(min, max)`: Requires between `min` and `max` arguments
- `NoArgs()`: No arguments allowed

### Flag Types

Cobra supports two types of flags:

1. **Persistent Flags**: Available to the command and all its subcommands
   ```go
   rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose output")
   ```

   Persistent flags are inherited by all subcommands. In the example above, the `--verbose` flag can be used with any
   subcommand:
   ```bash
   $ ./hello greet --verbose
   $ ./hello echo --verbose Hello
   ```

2. **Local Flags**: Only available to the specific command
   ```go
   greetCmd.Flags().StringVarP(&name, "name", "n", "World", "Name to greet")
   ```

   Local flags are only available to the command they're defined on. In the example above, the `--name` flag can only be
   used with the `greet` command:
   ```bash
   $ ./hello greet --name Alice  # Works
   $ ./hello echo --name Alice   # Error: unknown flag: --name
   ```

### Command Aliases

Make commands more user-friendly with aliases:

```go
greetCmd := &cobra.Command{
Use:     "greet",
Aliases: []string{"hello", "hi"},
// ...
}
```

Command aliases provide alternative names for your commands:

- **Aliases**: An array of strings that serve as alternative names for the command
- **Use case**: Helpful when users might expect different names for the same command
- **Example**: A user can type `git blame` or `git who` (an alias) to get the same functionality

Users can then run any of these commands:

```bash
$ ./hello greet --name Alice
$ ./hello hello --name Alice  # Same as "greet"
$ ./hello hi --name Alice     # Same as "greet"
```

## Resources

- [Official Cobra Documentation](https://github.com/spf13/cobra)
- [Cobra User Guide](https://github.com/spf13/cobra/blob/master/user_guide.md)
