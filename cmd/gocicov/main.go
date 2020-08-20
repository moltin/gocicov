package main

import (
	"fmt"
	"os"

	"github.com/moltin/gocicov/internal/coverage"
	"github.com/moltin/gocicov/internal/forcetest"
)

func main() {
	forceTest := forcetest.New(".")
	forceTest.Prepare()
	cov, err := coverage.Get()
	forceTest.Cleanup()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	fmt.Fprintf(os.Stderr, "Total coverage: %s\n", cov)
	base := coverage.FromEnv("COVERAGE")
	threshold := coverage.FromEnv("COVERAGE_THRESHOLD")
	diff := cov - base
	fmt.Fprintf(os.Stderr, "Coverage diff: %s\n", diff)
	if threshold+diff < 0 {
		fmt.Fprintf(os.Stderr, "Coverage diff failed to reach threshold of %s\n", threshold)
		os.Exit(1)
	}
	fmt.Println(cov.Int())
}
