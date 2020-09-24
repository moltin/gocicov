package modules

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type Module struct {
	ImportPath string
	Path       string
}

type List []Module

func FromFile(file string) (List, error) {
	data, err := ioutil.ReadFile(file)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	paths := []string{}
	for _, line := range strings.Split(string(data), "\n") {
		path := strings.TrimSpace(line)
		if len(path) == 0 || strings.HasPrefix(path, "#") {
			continue
		}
		paths = append(paths, path)
	}
	return Load(paths...)
}

func Load(paths ...string) (List, error) {
	args := []string{"list", "-f", "{{.Dir}} {{.ImportPath}}"}
	args = append(args, paths...)
	cmd := exec.Command("go", args...)
	data, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	mods := make([]Module, 0, len(lines))
	for _, line := range lines {
		words := strings.Split(line, " ")
		if len(words) != 2 {
			continue
		}
		mods = append(mods, Module{
			ImportPath: words[1],
			Path:       words[0],
		})
	}
	return List(mods), nil
}

func (m List) Filter(f List) List {
	r := make([]Module, 0, len(m))
	for _, mod := range m {
		if f.HasImport(mod.ImportPath) {
			continue
		}
		r = append(r, mod)
	}
	return List(r)
}

func (m List) HasImport(path string) bool {
	for _, i := range m.Imports() {
		if path == i {
			return true
		}
	}
	return false
}

func (m List) Imports() []string {
	r := make([]string, len(m))
	for i, mod := range m {
		r[i] = mod.ImportPath
	}
	return r
}

func (m List) Paths() []string {
	r := make([]string, len(m))
	for i, mod := range m {
		r[i] = mod.Path
	}
	return r
}
