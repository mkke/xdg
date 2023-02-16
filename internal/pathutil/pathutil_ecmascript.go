//go:build js && ecmascript
// +build js,ecmascript

package pathutil

// Exists returns true if the specified path exists.
func Exists(path string) bool {
	return false
}

// ExpandHome substitutes `~` and `$HOME` at the start of the specified
// `path` using the provided `home` location.
func ExpandHome(path, home string) string {
	return path
}
