package getarchitecture

import "runtime"

func GetArchitecture() string {
	arch := runtime.GOOS
	if arch == "amd64" {
		return "x86_64"
	}
	return "any"
}
