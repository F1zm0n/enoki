package pkgremover

import "os"

func RemovePkgs(pkgs []string, path string) error {
	err := os.Chdir(path)
	if err != nil {
		return err
	}

	for _, pkg := range pkgs {
		err := os.Remove(pkg)
		if err != nil {
			return err
		}
	}
	return nil
}
