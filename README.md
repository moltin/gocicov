# gocicov

`gocicov` is a tool for generating go code coverage inside a CI pipeline.  It
runs the tests, writing the test output and a code coverage summary to standard
error. It then returns an integer value for the total code coverage to standard
output which may be used to track code coverage. This integer is the total code
coverage percentage multplied by 100.

Unlike normal code coverage with `go test` the total percentage includes
packages for which no test files exist.

## Example usage

```
$ coverage=$(go run github.com/moltin/gocicov/cmd/gocicov)
testing: warning: no tests to run
PASS
coverage: 0.0% of statements
ok  	github.com/moltin/gocicov/cmd/gocicov	0.286s	coverage: 0.0% of statements [no tests to run]
testing: warning: no tests to run
PASS
coverage: 0.0% of statements
ok  	github.com/moltin/gocicov/internal/coverage	0.152s	coverage: 0.0% of statements [no tests to run]
testing: warning: no tests to run
PASS
coverage: 0.0% of statements
ok  	github.com/moltin/gocicov/internal/forcetest	0.411s	coverage: 0.0% of statements [no tests to run]
Total coverage: 0.000000
$ echo $coverage
0
```
