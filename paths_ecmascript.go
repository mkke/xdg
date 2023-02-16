//go:build js && ecmascript
// +build js,ecmascript

package xdg

func homeDir() string {
	return "/"
}

func initDirs(_ string) {
}

func initBaseDirs(_ string) {
}

func initUserDirs(_ string) {
}
