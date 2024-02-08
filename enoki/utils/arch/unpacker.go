package arch

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/klauspost/compress/zstd"
)

func UnpakAndInstPacman(
	repos []string,
	archName, pkgname, dir string,
) ([]string, error) {
	arr := make([]string, 0)
	data, err := GetFromPacmanBoth(repos, archName, pkgname)
	if err != nil {
		return nil, err
	}

	err = os.Chdir(dir)
	if err != nil {
		return nil, err
	}

	reader, err := zstd.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	// Create a tar reader from the zstd reader
	tr := tar.NewReader(reader)

	// Iterate through the files in the tar archive
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		fmt.Println(pkgname + "nigga" + archName)
		if err != nil {
			return nil, err
		}
		fmt.Println(header.Name)
		arr = append(arr, header.Name)

		if strings.Contains(header.Name, "/") {
			continue
		}

		// Create the file
		targetFile, err := os.Create(header.Name)
		if err != nil {
			return nil, err
		}
		defer targetFile.Close()

		// Write the file content
		if _, err := io.Copy(targetFile, tr); err != nil {
			return nil, err
		}
	}

	fmt.Println("installed ", pkgname)
	return arr, nil
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
