package env

import "runtime"

func Version() string {
	return runtime.Version()
}
