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
	defer forceTest.Cleanup()
	coverage, err := coverage.Coverage()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "Total coverage: %f\n", coverage)
	fmt.Printf("%d\n", int(coverage*100))
}
