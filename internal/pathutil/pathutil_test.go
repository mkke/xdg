package pathutil_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/adrg/xdg/internal/pathutil"
	"github.com/stretchr/testify/require"
)

func TestExists(t *testing.T) {
	tempDir := os.TempDir()

	// Test regular file.
	pathFile := filepath.Join(tempDir, "regular")
	f, err := os.Create(pathFile)
	require.NoError(t, err)
	require.NoError(t, f.Close())
	require.True(t, pathutil.Exists(pathFile))

	// Test symlink.
	pathSymlink := filepath.Join(tempDir, "symlink")
	require.NoError(t, os.Symlink(pathFile, pathSymlink))
	require.True(t, pathutil.Exists(pathSymlink))

	// Test non-existent file.
	require.NoError(t, os.Remove(pathFile))
	require.False(t, pathutil.Exists(pathFile))
	require.False(t, pathutil.Exists(pathSymlink))
	require.NoError(t, os.Remove(pathSymlink))
	require.False(t, pathutil.Exists(pathSymlink))
}

func TestCreate(t *testing.T) {
	tempDir := os.TempDir()

	// Test path selection order.
	p, err := pathutil.Create("test", []string{tempDir, "\000a"})
	require.NoError(t, err)
	require.Equal(t, filepath.Join(tempDir, "test"), p)

	p, err = pathutil.Create("test", []string{"\000a", tempDir})
	require.NoError(t, err)
	require.Equal(t, filepath.Join(tempDir, "test"), p)

	// Test relative parent directories.
	expected := filepath.Join(tempDir, "appname", "config", "test")
	p, err = pathutil.Create(filepath.Join("appname", "config", "test"), []string{"\000a", tempDir})
	require.NoError(t, err)
	require.Equal(t, expected, p)
	t.Log(expected)
	require.NoError(t, os.RemoveAll(filepath.Dir(expected)))

	expected = filepath.Join(tempDir, "appname", "test")
	p, err = pathutil.Create(filepath.Join("appname", "test"), []string{"\000a", tempDir})
	require.NoError(t, err)
	require.Equal(t, expected, p)
	t.Log(expected)
	require.NoError(t, os.RemoveAll(filepath.Dir(expected)))

	// Test invalid paths.
	_, err = pathutil.Create(filepath.Join("appname", "test"), []string{"\000a"})
	require.Error(t, err)

	_, err = pathutil.Create("test", []string{filepath.Join(tempDir, "\000a")})
	require.Error(t, err)
}
