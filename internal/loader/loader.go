// Package loader is the filesystem sheet.Loader for embedded sub-sheets: it
// resolves a SHEET("path") reference to a parsed sheet, confined to a root
// directory via os.Root so a reference can never escape it (no `..` traversal,
// no symlink escape). The engine (internal/sheet) stays filesystem-free; this
// package is injected by the frontends (serve, cli).
package loader

import (
	"io"
	"os"
	"path/filepath"

	"github.com/uplang/tsvsheet.go/internal/sheet"
)

// Dir is a directory path that bounds sheet resolution: every SHEET reference
// resolves within it.
type Dir string

// FS returns a sheet.Loader confined to root. A reference is resolved relative to
// the embedding sheet's own directory and opened through os.Root, which rejects
// any path escaping root (no `..` traversal, no symlink escape). The root is
// opened per call, so an unopenable root surfaces as a load error (#REF!) rather
// than failing construction.
func FS(root Dir) sheet.Loader {
	return func(base, ref sheet.Path) (sheet.Sheet, sheet.Path, error) {
		confined, err := os.OpenRoot(string(root))
		if err != nil {
			return sheet.Sheet{}, "", err
		}
		defer func() { _ = confined.Close() }()
		target := filepath.Clean(filepath.Join(filepath.Dir(string(base)), string(ref)))
		return load(confined, sheet.Path(target))
	}
}

// load opens target within the confined root, reads it, and parses it. Its
// resolved path (root-relative) is returned for cycle detection and as the base
// for the sub-sheet's own references.
func load(confined *os.Root, target sheet.Path) (sheet.Sheet, sheet.Path, error) {
	file, err := confined.Open(string(target))
	if err != nil {
		return sheet.Sheet{}, "", err
	}
	defer func() { _ = file.Close() }()
	data, err := io.ReadAll(file)
	if err != nil {
		return sheet.Sheet{}, "", err
	}
	parsed, err := sheet.Parse(data)
	return parsed, target, err
}
