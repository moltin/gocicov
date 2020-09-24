package coverage

import (
	"errors"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/moltin/gocicov/internal/modules"
)

func filterModules(modules modules.List) []string {
	r := make([]string, 0, len(modules))
	for _, mod := range modules {
		file := mod.Path + "/.skip_coverage"
		_, err := os.Stat(file)
		if err == nil {
			continue
		}
		r = append(r, mod.ImportPath)
	}
	return r
}

func runTests(modules modules.List) error {
	args := []string{"test", "-coverpkg=./...", "-v", "-coverprofile", "c.out"}
	args = append(args, filterModules(modules)...)
	cmd := exec.Command("go", args...)
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

func Get(modules modules.List) (Coverage, error) {
	err := runTests(modules)
	if err != nil {
		return 0, err
	}
	coverage, err := gatherCoverage()
	if err != nil {
		return 0, err
	}
	return parseCoverage(coverage)
}
