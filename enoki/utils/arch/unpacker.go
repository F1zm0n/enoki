package arch

import (
	"archive/tar"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/klauspost/compress/zstd"

	"github.com/F1zm0n/enoki/enoki/utils/entity"
)

func (a *PacmanApp) UnpakAndInstPacman(info entity.ArchInfo) ([]string, error) {
	fmt.Println(info.PkgName)
	arr := make([]string, 0)
	data, err := a.GetFromPacman(info)
	if err != nil {
		return nil, err
	}
	fmt.Println("this pkg: ", info.PkgName)

	dir := a.Path + "/" + info.PkgName

	err = os.MkdirAll(dir, 0770)

	if !errors.Is(err, os.ErrExist) {
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
		fmt.Println(header.Name)
		if err == io.EOF {
			break
		}
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

	fmt.Println("installed ", info.PkgName)
	return arr, nil
}
