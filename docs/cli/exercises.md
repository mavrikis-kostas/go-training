# CLI Exercises with Cobra

This section contains practical exercises to help you apply what you've learned about building command-line interfaces with Cobra.

## :material-code-braces-box: Exercise 1: Adding Flags to a Command

In this exercise, you'll practice adding flags to an existing command.

### Given Implementation

The following code is already implemented in a file named `main.go`:

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
        Use:   "greet",
        Short: "A simple greeting CLI",
        Long:  "A simple CLI application that greets the user",
        Run: func(cmd *cobra.Command, args []string) {
            // This will be modified to use the flags you add
            fmt.Println("Hello, World!")
        },
    }

    // TODO: Add flags here

    // Execute the root command
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
```

### Requirements

1. Add a string flag named `name` with a short form `n` and default value "World"
2. Add a boolean flag named `uppercase` with a short form `u` and default value false
3. Modify the Run function to use these flags:
   - Greet the user with the name provided
   - If uppercase is true, convert the greeting to uppercase

### Expected Output

When running with default flags:
```
Hello, World!
```

When running with custom name:
```
Hello, Alice!
```

When running with uppercase flag:
```
HELLO, WORLD!
```

When running with both flags:
```
HELLO, ALICE!
```

### Hints

1. Use `rootCmd.Flags().StringVarP()` to add a string flag
2. Use `rootCmd.Flags().BoolVarP()` to add a boolean flag
3. Use `cmd.Flags().GetString()` and `cmd.Flags().GetBool()` to retrieve flag values
4. Use `strings.ToUpper()` to convert a string to uppercase

??? example "Click for solution"

    Here's a complete solution to the exercise:

    ```go
    package main

    import (
        "fmt"
        "os"
        "strings"

        "github.com/spf13/cobra"
    )

    func main() {
        var name string
        var uppercase bool

        // Create the root command
        rootCmd := &cobra.Command{
            Use:   "greet",
            Short: "A simple greeting CLI",
            Long:  "A simple CLI application that greets the user",
            Run: func(cmd *cobra.Command, args []string) {
                greeting := fmt.Sprintf("Hello, %s!", name)

                if uppercase {
                    greeting = strings.ToUpper(greeting)
                }

                fmt.Println(greeting)
            },
        }

        // Add flags to the root command
        rootCmd.Flags().StringVarP(&name, "name", "n", "World", "Name to greet")
        rootCmd.Flags().BoolVarP(&uppercase, "uppercase", "u", false, "Convert greeting to uppercase")

        // Execute the root command
        if err := rootCmd.Execute(); err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
    }
    ```

## :material-code-braces-box: Exercise 2: Creating a Subcommand

In this exercise, you'll practice adding a subcommand to an existing CLI application.

### Given Implementation

The following code is already implemented in a file named `main.go`:

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
        Use:   "app",
        Short: "A sample CLI application",
        Long:  "A sample CLI application with subcommands",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("Use --help to see available commands")
        },
    }

    // TODO: Add your subcommand here

    // Execute the root command
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
```

### Requirements

1. Create a subcommand named `echo` that echoes the arguments passed to it
2. The subcommand should take at least one argument
3. Add a flag to the subcommand that allows reversing the output

### Expected Output

When running the echo command:
```
$ ./app echo hello world
hello world
```

When running with the reverse flag:
```
$ ./app echo --reverse hello world
dlrow olleh
```

### Hints

1. Create a function that returns a `*cobra.Command` for your subcommand
2. Use `cobra.MinimumNArgs(1)` to require at least one argument
3. Use `strings.Join(args, " ")` to combine all arguments
4. To reverse a string, you can use a simple loop or create a helper function

??? example "Click for solution"

    Here's a complete solution to the exercise:

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
            Use:   "app",
            Short: "A sample CLI application",
            Long:  "A sample CLI application with subcommands",
            Run: func(cmd *cobra.Command, args []string) {
                fmt.Println("Use --help to see available commands")
            },
        }

        // Add the echo subcommand
        rootCmd.AddCommand(echoCmd())

        // Execute the root command
        if err := rootCmd.Execute(); err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
    }

    func echoCmd() *cobra.Command {
        var reverse bool

        cmd := &cobra.Command{
            Use:   "echo [string to echo]",
            Short: "Echo a string",
            Args:  cobra.MinimumNArgs(1),
            Run: func(cmd *cobra.Command, args []string) {
                // Join all arguments
                message := strings.Join(args, " ")

                // Reverse the message if the flag is set
                if reverse {
                    message = reverseString(message)
                }

                fmt.Println(message)
            },
        }

        // Add flags
        cmd.Flags().BoolVarP(&reverse, "reverse", "r", false, "Reverse the output")

        return cmd
    }

    // Helper function to reverse a string
    func reverseString(s string) string {
        runes := []rune(s)
        for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
            runes[i], runes[j] = runes[j], runes[i]
        }
        return string(runes)
    }
    ```

## :material-code-braces-box: Exercise 3: Implementing a Calculator Subcommand

In this exercise, you'll implement a calculator subcommand for a CLI application.

### Given Implementation

The following code is already implemented in a file named `main.go`:

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
        Use:   "calc",
        Short: "A simple calculator CLI",
        Long:  "A simple CLI application that performs basic calculations",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("Use --help to see available commands")
        },
    }

    // TODO: Implement and add your calculator subcommand here

    // Execute the root command
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
```

### Requirements

1. Create a subcommand named `add` that adds two numbers provided as arguments
2. The subcommand should:
   - Take exactly two arguments (numbers to add)
   - Convert string arguments to integers
   - Print the result of the addition
3. Add a flag to the subcommand that allows displaying the calculation in verbose mode

### Expected Output

When running the add command:
```
$ ./calc add 5 3
8
```

When running with the verbose flag:
```
$ ./calc add --verbose 5 3
5 + 3 = 8
```

### Hints

1. Use `cobra.ExactArgs(2)` to require exactly two arguments
2. Use `strconv.Atoi()` to convert string arguments to integers
3. Handle conversion errors appropriately
4. Use a boolean flag for the verbose mode

??? example "Click for solution"

    Here's a complete solution to the exercise:

    ```go
    package main

    import (
        "fmt"
        "os"
        "strconv"

        "github.com/spf13/cobra"
    )

    func main() {
        // Create the root command
        rootCmd := &cobra.Command{
            Use:   "calc",
            Short: "A simple calculator CLI",
            Long:  "A simple CLI application that performs basic calculations",
            Run: func(cmd *cobra.Command, args []string) {
                fmt.Println("Use --help to see available commands")
            },
        }

        // Add the calculator subcommand
        rootCmd.AddCommand(addCmd())

        // Execute the root command
        if err := rootCmd.Execute(); err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
    }

    func addCmd() *cobra.Command {
        var verbose bool

        cmd := &cobra.Command{
            Use:   "add [number1] [number2]",
            Short: "Add two numbers",
            Args:  cobra.ExactArgs(2),
            Run: func(cmd *cobra.Command, args []string) {
                // Convert arguments to integers
                num1, err := strconv.Atoi(args[0])
                if err != nil {
                    fmt.Printf("Error: %s is not a valid number\n", args[0])
                    return
                }

                num2, err := strconv.Atoi(args[1])
                if err != nil {
                    fmt.Printf("Error: %s is not a valid number\n", args[1])
                    return
                }

                // Calculate the sum
                sum := num1 + num2

                // Display the result
                if verbose {
                    fmt.Printf("%d + %d = %d\n", num1, num2, sum)
                } else {
                    fmt.Println(sum)
                }
            },
        }

        // Add flags
        cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Display calculation details")

        return cmd
    }
    ```
