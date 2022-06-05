package utils

import "runtime"

func IsUnix() bool {
	return TextContainOne(runtime.GOOS, "linux", "darwin", "netbsd", "freebsd")
}
