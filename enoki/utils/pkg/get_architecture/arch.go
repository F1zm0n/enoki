package getarchitecture

func GetArchitecture(arch string) string {
	if arch == "amd64" {
		return "x86_64"
	}
	return "any"
}
