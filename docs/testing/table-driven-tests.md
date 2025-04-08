# Table-Driven Tests in Go

Table-driven tests are a common pattern in Go for testing multiple inputs and expected outputs in a more organized and
maintainable way.

## Why Use Table-Driven Tests?

When you have multiple test cases for the same function, table-driven tests offer several advantages:

1. **Reduced Repetition**: You don't need to write similar test code for each test case
2. **Better Organization**: All test cases are grouped together in a clear structure
3. **Easier Maintenance**: Adding new test cases is as simple as adding a new entry to the table
4. **Improved Readability**: The test logic is separated from the test data

## Example: Testing the Add Function

Let's use the same `Add` function from the [Basic Unit Testing](basic-unit-tests.md) example:

```go
package math

// Add returns the sum of two integers
func Add(a, b int) int {
	return a + b
}
```

## Writing a Table-Driven Test

Here's how to write a table-driven test for the `Add` function:

```go
package math

import (
	"testing"
)

func TestAddTable(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -2, -3, -5},
		{"mixed numbers", -2, 5, 3},
		{"zeros", 0, 0, 0},
	}

	// Run all test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Add(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("Add(%d, %d) = %d; expected %d",
					tc.a, tc.b, result, tc.expected)
			}
		})
	}
}
```

## Understanding Table-Driven Tests

Let's break down the key components:

1. **Test Case Structure**: We define a struct that contains all the information needed for each test case:
    - A name for the test case
    - Input values (a and b)
    - Expected output

2. **Test Cases Table**: We create a slice of these structs, with each entry representing a test case.

3. **Iterating Through Test Cases**: We loop through each test case and run the test.

4. **Subtests**: We use `t.Run()` to create a subtest for each test case, which:
    - Improves test output organization
    - Allows running specific test cases using the `-run` flag
    - Provides better failure reporting

## Benefits of This Approach

With this approach:

- Adding a new test case is as simple as adding a new entry to the table
- The test logic is written only once, reducing the chance of errors
- Test failures clearly indicate which specific test case failed
- The code is more maintainable and easier to understand

## Next Steps

Now that you understand table-driven tests, you might want to explore:

- [Running Tests](running-tests.md) to learn how to execute your tests and interpret the results
- Return to the [Testing Overview](index.md) for more testing topics