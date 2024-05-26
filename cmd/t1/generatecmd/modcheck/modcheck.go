package modcheck

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/senforsce/t1"

	"golang.org/x/mod/modfile"
	"golang.org/x/mod/semver"
)

// WalkUp the directory tree, starting at dir, until we find a directory containing
// a go.mod file.
func WalkUp(dir string) (string, error) {
	dir, err := filepath.Abs(dir)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path: %w", err)
	}

	var modFile string
	for {
		modFile = filepath.Join(dir, "go.mod")
		_, err := os.Stat(modFile)
		if err != nil && !os.IsNotExist(err) {
			return "", fmt.Errorf("failed to stat go.mod file: %w", err)
		}
		if os.IsNotExist(err) {
			// Move up.
			prev := dir
			dir = filepath.Dir(dir)
			if dir == prev {
				break
			}
			continue
		}
		break
	}

	// No file found.
	if modFile == "" {
		return dir, fmt.Errorf("could not find go.mod file")
	}
	return dir, nil
}

// Replace "go 1.21.3" with "go 1.21" until https://github.com/golang/go/issues/61888 is fixed, see t1 issue https://github.com/senforsce/t1/issues/355
var goVersionRegexp = regexp.MustCompile(`\ngo (\d+\.\d+)(?:\D.+)\n`)

func patchGoVersion(moduleFileContents []byte) []byte {
	return goVersionRegexp.ReplaceAll(moduleFileContents, []byte("\ngo $1\n"))
}

func Check(dir string) error {
	dir, err := WalkUp(dir)
	if err != nil {
		return err
	}

	// Found a go.mod file.
	// Read it and find the t1 version.
	modFile := filepath.Join(dir, "go.mod")
	m, err := os.ReadFile(modFile)
	if err != nil {
		return fmt.Errorf("failed to read go.mod file: %w", err)
	}

	// Replace "go 1.21.x" with "go 1.21".
	m = patchGoVersion(m)

	mf, err := modfile.Parse(modFile, m, nil)
	if err != nil {
		return fmt.Errorf("failed to parse go.mod file: %w", err)
	}
	if mf.Module.Mod.Path == "github.com/senforsce/t1" {
		// The go.mod file is for t1 itself.
		return nil
	}
	for _, r := range mf.Require {
		if r.Mod.Path == "github.com/senforsce/t1" {
			cmp := semver.Compare(r.Mod.Version, t1.Version())
			if cmp < 0 {
				return fmt.Errorf("generator %v is newer than t1 version %v found in go.mod file, consider running `go get -u github.com/senforsce/t1` to upgrade", t1.Version(), r.Mod.Version)
			}
			if cmp > 0 {
				return fmt.Errorf("generator %v is older than t1 version %v found in go.mod file, consider upgrading t1 CLI", t1.Version(), r.Mod.Version)
			}
			return nil
		}
	}
	return fmt.Errorf("t1 not found in go.mod file, run `go get github.com/senforsce/t1 to install it`")
}
