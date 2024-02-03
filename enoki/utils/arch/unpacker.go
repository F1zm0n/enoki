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

func UnpakAndInstPacman(archName, pkgname, dir string) error {
	data, err := GetFromPacmanBoth(archName, pkgname)
	if err != nil {
		return err
	}

	reader, err := zstd.NewReader(bytes.NewReader(data))
	if err != nil {
		return err
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
		if err != nil {
			panic(err)
		}
		fmt.Println(header.Name) // todo убрать потом

		if strings.Contains(header.Name, "/") {
			continue
		}

		// Create the file
		targetFile, err := os.Create(header.Name)
		if err != nil {
			return err
		}
		defer targetFile.Close()

		// Write the file content
		if _, err := io.Copy(targetFile, tr); err != nil {
			return err
		}
	}

	return nil
}
