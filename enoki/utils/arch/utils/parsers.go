package pacutils

import (
	"bufio"
	"bytes"
	"strings"
)

// ParseMirrors parses mirror links and return slice of links
// to get data from
func ParseMirrors(data []byte) ([]string, error) {
	f := bytes.NewReader(data)

	arr := make([]string, 0)

	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "##") {
			continue
		}
		mirror := strings.Trim(scanner.Text(), "#Server =")
		mirror = strings.TrimSpace(mirror)
		arr = append(arr, mirror)
	}

	return arr, nil
}
