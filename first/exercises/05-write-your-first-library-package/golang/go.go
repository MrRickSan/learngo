package golang

import "runtime"

// Version retuns the installed go version
func Version() string {
	return runtime.Version()
}
