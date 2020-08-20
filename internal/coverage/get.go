package coverage

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func runTests() error {
	cmd := exec.Command("go", "test", "-v", "-coverpkg=./...", "-coverprofile=c.out", "./...")
	cmd.Stdout = os.Stderr
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func gatherCoverage() (string, error) {
	cmd := exec.Command("go", "tool", "cover", "-func=c.out")
	data, err := cmd.Output()
	return string(data), err
}

func parseCoverage(coverage string) (Coverage, error) {
	fmt.Fprintf(os.Stderr, "%s\n", coverage)
	lines := strings.Split(coverage, "\n")
	if len(lines) == 0 {
		return 0, errors.New("no coverage data")
	}
	total := ""
	for _, line := range lines {
		if strings.HasPrefix(line, "total") {
			total = line
			break
		}
	}
	if total == "" {
		return 0, errors.New("no coverage data")
	}

	words := strings.Fields(total)
	if len(words) < 3 {
		return 0, errors.New("malformed coverage data")
	}
	percent := strings.TrimSuffix(words[len(words)-1], "%")

	cov, err := strconv.ParseFloat(percent, 64)
	return Coverage(cov), err
}

func Get() (Coverage, error) {
	defer os.Remove("c.out")
	err := runTests()
	if err != nil {
		return 0, nil
	}
	coverage, err := gatherCoverage()
	if err != nil {
		return 0, nil
	}
	return parseCoverage(coverage)
}
