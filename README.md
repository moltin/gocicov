# gocicov

`gocicov` is a tool for generating go code coverage inside a CI pipeline.  It
runs the tests, writing the test output and a code coverage summary to standard
error. It then returns an integer value for the total code coverage to standard
output which may be used to track code coverage. This integer is the total code
coverage in hundredths of a percent.

Unlike normal code coverage with `go test` the total percentage includes
packages for which no test files exist.

`gocicov` uses a pair of environment variables to enforce changes in code
coverage. `COVERAGE` is the existing coverage in hundredths of a percent.
`COVERAGE_THRESHOLD` is the maximum permissable drop in code coverage,
also in hundredths of a percent.

`gocicov` fails if any of the tests fail or if the tests pass but the coverage threshold is not met.

## Example usage

```
$ newcoverage=$(COVERAGE=100 COVERAGE_THRESHOLD=200 gocicov)
WARNING: tests needed for cmd/gocicov
WARNING: tests needed for internal/coverage
WARNING: tests needed for internal/forcetest
testing: warning: no tests to run
PASS
coverage: 0.0% of statements in ./...
ok  	github.com/moltin/gocicov/cmd/gocicov	0.180s	coverage: 0.0% of statements in ./... [no tests to run]
testing: warning: no tests to run
PASS
coverage: 0.0% of statements in ./...
ok  	github.com/moltin/gocicov/internal/coverage	0.484s	coverage: 0.0% of statements in ./... [no tests to run]
testing: warning: no tests to run
PASS
coverage: 0.0% of statements in ./...
ok  	github.com/moltin/gocicov/internal/forcetest	0.323s	coverage: 0.0% of statements in ./... [no tests to run]
Total coverage: 0.00%
Coverage diff: -1.00%
$ echo $newcoverage
0
```

## Example buddy action

Having set up project-level variables `COVERAGE` (should be settable) and
`COVERAGE_THRESHOLD`:
```
  actions:
    - action: "Unit Tests and Linting"
      type: "BUILD"
      docker_image_name: "library/golang"
      docker_image_tag: "1.14-stretch"
      setup_commands:
        - "go get -u github.com/moltin/gocicov/cmd/gocicov"
      execute_commands:
        - "newcoverage=$(gocicov)"
        - "if [ $BUDDY_EXECUTION_BRANCH = master ]; then COVERAGE=$newcoverage ; fi"
```
