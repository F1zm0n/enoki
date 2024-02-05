package pkgremover

import (
	"io/fs"
	"path/filepath"
)

func GetPkgToInst(startPath string, deps []string) ([]string, error) {
	for i, subDirToSkip := range deps {
		err := filepath.Walk(startPath, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() && info.Name() == subDirToSkip {
				deps = remove(deps, i)

				return filepath.SkipDir
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	return deps, nil
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
