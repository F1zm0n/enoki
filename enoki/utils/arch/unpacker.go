package arch

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/klauspost/compress/zstd"
)

func UnpakAndInstPacman(
	timeout time.Duration,
	arr []string,
	archName, pkgname, dir string,
) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	data, err := GetFromPacmanBoth(ctx, archName, pkgname)
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
		fmt.Println(pkgname, archName)
		if err != nil {
			panic(err)
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
