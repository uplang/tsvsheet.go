package loader_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/uplang/tsvsheet.go/internal/loader"
	"github.com/uplang/tsvsheet.go/internal/sheet"
)

// write creates a file under dir.
func write(t *testing.T, dir, name, content string) {
	t.Helper()
	require.NoError(t, os.WriteFile(filepath.Join(dir, name), []byte(content), 0o600))
}

func TestFS_OpenRootFails(t *testing.T) {
	t.Parallel()

	// A root that cannot be opened surfaces as a load error, not a build error.
	ld := loader.FS(loader.Dir(filepath.Join(t.TempDir(), "does-not-exist")))
	_, _, err := ld("main.tsvt", "child.tsvt")
	require.Error(t, err)
}

func TestFS_LoadsAndResolvesRelative(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	write(t, dir, "child.tsvt", "=output(42)\n")
	ld := loader.FS(loader.Dir(dir))

	sub, resolved, err := ld("main.tsvt", "child.tsvt")
	require.NoError(t, err)
	assert.Equal(t, sheet.Path("child.tsvt"), resolved)
	assert.Equal(t, "42", sub.Compute()[0][0])
}

func TestFS_EscapeIsRejected(t *testing.T) {
	t.Parallel()

	ld := loader.FS(loader.Dir(t.TempDir()))

	_, _, err := ld("main.tsvt", "../escape.tsvt")
	require.Error(t, err) // os.Root refuses to leave the root
}

func TestFS_MissingFile(t *testing.T) {
	t.Parallel()

	ld := loader.FS(loader.Dir(t.TempDir()))

	_, _, err := ld("main.tsvt", "absent.tsvt")
	require.Error(t, err)
}

func TestFS_ReadAllErrorOnDirectory(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	require.NoError(t, os.Mkdir(filepath.Join(dir, "nested"), 0o750))
	ld := loader.FS(loader.Dir(dir))

	_, _, err := ld("main.tsvt", "nested") // opens, but reading a directory fails
	require.Error(t, err)
}

func TestFS_ParseError(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	write(t, dir, "bad.tsvt", "=sum(\n") // malformed formula
	ld := loader.FS(loader.Dir(dir))

	_, _, err := ld("main.tsvt", "bad.tsvt")
	require.Error(t, err)
}
