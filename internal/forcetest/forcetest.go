package forcetest

import (
	"fmt"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Forcer struct {
	dir   string
	files []string
}

func New(dir string) *Forcer {
	if dir == "" {
		dir = "."
	}
	return &Forcer{dir: dir}
}

func getPackageName(file string) (string, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, file, nil, parser.PackageClauseOnly)
	if err != nil {
		return "", err
	}
	return f.Name.Name, nil
}

func (f *Forcer) addForcer(path, pkgnam string) error {
	file := path + "/force_coverage_test.go"
	content := fmt.Sprintf("package %s\n", pkgnam)
	err := ioutil.WriteFile(file, []byte(content), 0644)
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "WARNING: tests needed for %s\n", path)
	f.files = append(f.files, file)
	return nil
}

func (f *Forcer) checkDir(path string) error {
	if m, err := filepath.Glob(path + "/*_test.go"); err != nil || len(m) > 0 {
		return err
	}
	m, err := filepath.Glob(path + "/*.go")
	if err != nil || len(m) == 0 {
		return err
	}
	pkgnam, err := getPackageName(m[0])
	if err != nil {
		return err
	}
	return f.addForcer(path, pkgnam)
}

func (f *Forcer) Prepare() error {
	filepath.Walk(f.dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return f.checkDir(path)
		}
		return nil
	})
	return nil
}

func (f *Forcer) Cleanup() {
	for _, file := range f.files {
		err := os.Remove(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to delete %s: %v", file, err)
		}
	}
}
