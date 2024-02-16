package pacutils

import (
	"os"
	"strings"
)

func TrimSoPacman(s string) string {
	idx := strings.Index(s, ".so")
	if idx == -1 {
		return s
	}

	str := strings.Split(s, "")

	for i := idx; i < len(s); i++ {
		str[i] = " "
	}
	s = strings.Join(str, "")
	s = strings.TrimSpace(s)
	return s
}

func TrimVersionPacman(s string) string {
	idx := strings.Index(s, ">")
	if idx == -1 {
		idx = strings.Index(s, "<")
		if idx == -1 {
			return s
		}
		str := strings.Split(s, "")

		for i := idx; i < len(s); i++ {
			str[i] = " "
		}

		s = strings.Join(str, "")

		s = strings.TrimSpace(s)

		return s

	} else {

		str := strings.Split(s, "")

		for i := idx; i < len(s); i++ {
			str[i] = " "
		}
		s = strings.Join(str, "")
		s = strings.TrimSpace(s)
		return s
	}
}

func InstallPkgInDir(path string, pkg string) (*os.File, error) {
	if pkg == ".BUILDINFO" || pkg == ".MTREE" || pkg == ".PKGINFO" || pkg == ".INSTALL" {
		err := os.Chdir(path)
		if err != nil {
			return nil, err
		}
		f, err := os.Create(pkg)
		if err != nil {
			return nil, err
		}
		return f, nil
	}
	f, err := os.Create(pkg)
	if err != nil {
		return nil, err
	}
	return f, nil
}
