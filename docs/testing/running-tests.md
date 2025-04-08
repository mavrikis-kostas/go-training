# Running Tests in Go

This guide explains how to run tests in Go, interpret the results, and use various testing flags to enhance your testing
workflow.

## Basic Test Execution

To run all tests in the current directory:

```bash
go test
```

This command:

1. Looks for files with names ending in `_test.go`
2. Compiles the test files along with the code they're testing
3. Runs the tests
4. Reports the results

## Understanding Test Output

By default, `go test` only displays output for failed tests. A successful test run looks like:

```
PASS
ok      example/math    0.002s
```

If a test fails, you'll see output like:

```
--- FAIL: TestAdd (0.00s)
    math_test.go:15: Add(2, 3) = 6; expected 5
FAIL
exit status 1
FAIL    example/math    0.002s
```

## Useful Testing Flags

Go's testing tool provides several flags to customize test execution:

### Verbose Output

```bash
go test -v
```

The `-v` flag enables verbose output, showing each test that runs:

```
=== RUN   TestAdd
--- PASS: TestAdd (0.00s)
=== RUN   TestAddTable
=== RUN   TestAddTable/positive_numbers
=== RUN   TestAddTable/negative_numbers
=== RUN   TestAddTable/mixed_numbers
=== RUN   TestAddTable/zeros
--- PASS: TestAddTable (0.00s)
    --- PASS: TestAddTable/positive_numbers (0.00s)
    --- PASS: TestAddTable/negative_numbers (0.00s)
    --- PASS: TestAddTable/mixed_numbers (0.00s)
    --- PASS: TestAddTable/zeros (0.00s)
PASS
ok      example/math    0.002s
```

### Running Specific Tests

To run a specific test or set of tests, use the `-run` flag with a regular expression:

```bash
go test -run TestAdd        # Runs TestAdd and TestAddTable
go test -run TestAdd$       # Runs only TestAdd
go test -run TestAddTable   # Runs only TestAddTable
```

For table-driven tests with subtests, you can target specific subtests:

```bash
go test -run TestAddTable/positive  # Runs only the "positive numbers" subtest
```

### Code Coverage

To see how much of your code is covered by tests:

```bash
go test -cover
```

For a more detailed coverage report:

```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out  # Opens a browser with highlighted code coverage
```

## Testing Multiple Packages

To test all packages in your project:

```bash
go test ./...
```

To test a specific package:

```bash
go test example/math
```

## Best Practices for Go Testing

1. **Keep Tests Simple**: Tests should be easy to understand and maintain.

2. **Test One Thing at a Time**: Each test function should focus on testing one specific behavior.

3. **Use Table-Driven Tests**: For testing multiple inputs and expected outputs.

4. **Test Edge Cases**: Include tests for boundary conditions and error cases.

5. **Use Subtests for Organization**: Group related tests using `t.Run()`.

6. **Write Helpful Error Messages**: Error messages should clearly indicate what failed and why.

7. **Run Tests Regularly**: Integrate testing into your development workflow.

8. **Check Code Coverage**: Aim for high test coverage, but focus on testing critical paths.

## Next Steps

Now that you understand how to run tests in Go, you might want to explore:

- Return to [Basic Unit Testing](basic-unit-tests.md) or [Table-Driven Tests](table-driven-tests.md)
- Go back to the [Testing Overview](index.md) for more testing topics