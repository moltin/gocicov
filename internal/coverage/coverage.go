package coverage

import (
	"fmt"
	"os"
	"strconv"
)

type Coverage float64

func FromEnv(envVar string) Coverage {
	f, _ := strconv.ParseFloat(os.Getenv(envVar), 64)
	return Coverage(f / 100)
}

func (c Coverage) String() string {
	return fmt.Sprintf("%.2f%%", c)
}

func (c Coverage) Int() int {
	return int(c * 100)
}
