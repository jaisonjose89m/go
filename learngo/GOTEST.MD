# Unit Test
## Naming Conventions
It is important to follow these conventions as the test runner in go depend on these conventions.
1. Test file name `<*_test>`
2. Test functions should be named as `Test<meaningful_name>`
3. No built-in support for assertions. Have to use `if` for checking if the expectation is met.
4. White box testing - Same package name as that of the code under test
5. Black box testing - `<packagename_test>`

## Writing tests
### Immediate failure
Terminate the test function when this method is invoked
1. t.FailNow
2. t.Fatal
3. t.Fatalf

### Non-immediate failure
Execution continues even  after this method is invoked
1. t.Fail
2. t.Error
3. t.Errorf

### Example:
file name - main_test.go

```go
package main

import "testing"

func TestFoo(t *testing.T) {
	// test logic
}
```

## How to run test ?
1. `go test` - run test in current directory
2. `go test {pkg1} {pkg2} ... {pkgn}` - run test in specified packages
3. `go test ./...` - run test for current package and descendents
4. `go test -v` - run test in verbose mode. Can be used with any of the above options.
5. `go test -run {regex}` - Runs test only matching {regex}


## Get help & coverage
1. Test methods accept a `*testing.T` object that is injected by the test runner
2. `go help test` - basic help
3. `go help testflag` - comprehensive help
4. `go test -cover`
5. `go test -coverprofile cover.out ./...` & `go tool cover -func cover.out` or `go tool cover -html cover.out`

## Other useful functions
1. Log & Logf
2. Helper
3. Skip, Skipf & SkipNow
4. Run
5. Parallel

Note: all of these are available in `github.com/pkg/testing`

## Useful packages
1. testing -- basic package
2. testing/quick -- simplify black box testing
3. testing/iotest 
4. net/http/httptest

## Community Support 
1. Testify -- assertion framework
2. GinkGo -- behaviour driven testing
3. GoConvey -- helps to show results in browser based format
4. httpexpect
5. gomock
6. go-sqlmock by DATA-DOG

# Benchmark Test
## Naming Convention
file name - main_test.go

```go
package main_test

import "testing"

func BenchmarkFoo(b *testing.B) {
	// test logic
}
```

## How to run ?
1. `go test -bench`
2. `go test -bench benchtime 30s`
