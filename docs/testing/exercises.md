# Testing Exercises

This section contains practical exercises to help you apply what you've learned about testing in Go.

## :material-code-braces-box: Exercise 1: Basic Unit Testing

In this exercise, you'll practice writing basic unit tests for a simple function that is already implemented.

### Given Implementation

The following function is already implemented in a file named `calculator.go`:

```go
package calculator

// Multiply returns the product of two integers
func Multiply(a, b int) int {
    return a * b
}
```

### Requirements

1. Create a test file named `calculator_test.go` with a test function for `Multiply`
2. Test at least three different cases: positive numbers, negative numbers, and zeros

### Hints

1. Remember to use the same package name in both files
2. Use descriptive error messages in your test assertions
3. Follow Go's naming conventions for test functions (`TestMultiply`)

??? example "Click for solution"

    Here's a complete solution to the exercise:

    **calculator_test.go**:
    ```go
    package calculator

    import (
        "testing"
    )

    func TestMultiply(t *testing.T) {
        // Test case 1: positive numbers
        result := Multiply(3, 4)
        expected := 12
        if result != expected {
            t.Errorf("Multiply(3, 4) = %d; expected %d", result, expected)
        }

        // Test case 2: negative numbers
        result = Multiply(-2, -3)
        expected = 6
        if result != expected {
            t.Errorf("Multiply(-2, -3) = %d; expected %d", result, expected)
        }

        // Test case 3: mixed numbers
        result = Multiply(-5, 2)
        expected = -10
        if result != expected {
            t.Errorf("Multiply(-5, 2) = %d; expected %d", result, expected)
        }

        // Test case 4: zeros
        result = Multiply(0, 5)
        expected = 0
        if result != expected {
            t.Errorf("Multiply(0, 5) = %d; expected %d", result, expected)
        }
    }
    ```

## :material-code-braces-box: Exercise 2: Table-Driven Testing

In this exercise, you'll practice writing table-driven tests for a string manipulation function that is already implemented.

### Given Implementation

The following function is already implemented in a file named `strings.go`:

```go
package stringutils

// Reverse returns the reverse of the input string
func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
```

### Requirements

1. Create a test file named `strings_test.go` with a table-driven test for `Reverse`
2. Include at least five test cases, including empty string, palindromes, and strings with special characters

### Hints

1. Use a struct to define your test cases with input and expected output fields
2. Use `t.Run()` to create subtests for each test case
3. Give each test case a descriptive name

??? example "Click for solution"

    Here's a complete solution to the exercise:

    **strings_test.go**:
    ```go
    package stringutils

    import (
        "testing"
    )

    func TestReverse(t *testing.T) {
        // Define test cases
        testCases := []struct {
            name     string
            input    string
            expected string
        }{
            {"empty string", "", ""},
            {"single character", "a", "a"},
            {"palindrome", "radar", "radar"},
            {"simple string", "hello", "olleh"},
            {"with spaces", "hello world", "dlrow olleh"},
            {"with special characters", "hello, 世界!", "!界世 ,olleh"},
        }

        // Run all test cases
        for _, tc := range testCases {
            t.Run(tc.name, func(t *testing.T) {
                result := Reverse(tc.input)
                if result != tc.expected {
                    t.Errorf("Reverse(%q) = %q; expected %q", 
                        tc.input, result, tc.expected)
                }
            })
        }
    }
    ```

## :material-code-braces-box: Exercise 3: Testing with Coverage

In this exercise, you'll practice writing tests with good code coverage for a simple function that is already implemented.

### Given Implementation

The following function is already implemented in a file named `validator.go`:

```go
package validator

// IsInRange checks if a number is within the specified range (inclusive)
func IsInRange(num, min, max int) bool {
    // Check if the number is less than the minimum
    if num < min {
        return false
    }

    // Check if the number is greater than the maximum
    if num > max {
        return false
    }

    // If we get here, the number is within range
    return true
}
```

### Requirements

1. Create a test file named `validator_test.go` with tests for `IsInRange`
2. Ensure your tests cover various cases: numbers within range, below range, above range, and edge cases
3. Run your tests with coverage and aim for 100% coverage

### Hints

1. Use table-driven tests to handle multiple test cases efficiently
2. Don't forget to test the boundary values (exactly at min and max)
3. Use `go test -cover` to check your test coverage

??? example "Click for solution"

    Here's a complete solution to the exercise:

    **validator_test.go**:
    ```go
    package validator

    import (
        "testing"
    )

    func TestIsInRange(t *testing.T) {
        // Define test cases
        testCases := []struct {
            name     string
            num      int
            min      int
            max      int
            expected bool
        }{
            {"within range", 5, 1, 10, true},
            {"at minimum", 1, 1, 10, true},
            {"at maximum", 10, 1, 10, true},
            {"below range", 0, 1, 10, false},
            {"above range", 11, 1, 10, false},
            {"negative numbers", -5, -10, -1, true},
        }

        // Run all test cases
        for _, tc := range testCases {
            t.Run(tc.name, func(t *testing.T) {
                result := IsInRange(tc.num, tc.min, tc.max)
                if result != tc.expected {
                    t.Errorf("IsInRange(%d, %d, %d) = %v; expected %v", 
                        tc.num, tc.min, tc.max, result, tc.expected)
                }
            })
        }
    }
    ```

    To run the tests with coverage:
    ```bash
    go test -cover
    ```

    Expected output:
    ```
    PASS
    coverage: 100.0% of statements
    ok      validator       0.187s
    ```
