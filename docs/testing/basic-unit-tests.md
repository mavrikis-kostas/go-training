# Basic Unit Testing in Go

This guide demonstrates how to write simple unit tests in Go using the standard `testing` package.

## The Code to Test

Let's start with a simple function that we want to test. Save this code to a file named `math.go`:

```go
package math

// Add returns the sum of two integers
func Add(a, b int) int {
	return a + b
}
```

This is a very simple function, which makes it perfect for demonstrating testing concepts without getting distracted by
complex logic.

## Writing a Basic Test

Now, let's write a test for our `Add` function. Save this test code to a file named `math_test.go` in the same
directory:

```go
package math

import (
	"testing"
)

// TestAdd tests the Add function
func TestAdd(t *testing.T) {
	// Test case 1: positive numbers
	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2, 3) = %d; expected %d", result, expected)
	}

	// Test case 2: negative numbers
	result = Add(-2, -3)
	expected = -5

	if result != expected {
		t.Errorf("Add(-2, -3) = %d; expected %d", result, expected)
	}

	// Test case 3: mixed numbers
	result = Add(-2, 5)
	expected = 3

	if result != expected {
		t.Errorf("Add(-2, 5) = %d; expected %d", result, expected)
	}
}
```

## Understanding the Test

Let's break down the key components of our test:

1. **Package Declaration**: The test file uses the same package name as the code being tested.

2. **Import Statement**: We import the `testing` package, which provides the testing framework.

3. **Test Function**: Test functions must:
    - Start with the word `Test`
    - Take a pointer to `testing.T` as their only parameter
    - Be in the same package as the code they're testing

4. **Test Cases**: We test multiple scenarios:
    - Adding positive numbers
    - Adding negative numbers
    - Adding a mix of positive and negative numbers

5. **Assertions**: We compare the actual result with the expected result and report an error if they don't match.

## Next Steps

Now that you understand basic unit testing, you might want to explore:

- [Table-Driven Tests](table-driven-tests.md) for a more efficient way to organize multiple test cases
- [Running Tests](running-tests.md) to learn how to execute your tests and interpret the results