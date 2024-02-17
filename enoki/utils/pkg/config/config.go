package configs

import (
	"bufio"
	"os"
	"regexp"
	"strings"

	apperror "github.com/F1zm0n/enoki/enoki/utils/pkg/AppError"
)

type ConfigApp struct {
	Config      map[string]string
	MirrorLinks []string
}

// ReadConfig reads the config by the given path reads the map
// to the ConfigApp struct, if there are any mirror_link_{0-9}
// it will write in ConfigApp.MirrorLinks and will consider as a mirror link.
//
// So it will write to the map ConfigApp.Config and ConfigApp.MirrorLinks
// and will return an error apperror.ErrReadingConfig or err if something
// goes wrong.
func (c *ConfigApp) ReadConfig(path string) error {
	dat, err := os.ReadFile(path)
	if err != nil {
		return apperror.ErrReadingConfig
	}

	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			continue
		}

		str := strings.TrimSpace(t)

		if strings.Split(str, "")[0] == "#" {
			continue
		}

		if !strings.Contains(str, "=") {
			continue
		}

		arr := strings.Split(str, "=")

		k := strings.TrimSpace(arr[0])
		v := strings.TrimSpace(arr[1])

		reg, err := regexp.Compile(`mirror_link_\d`)
		if err != nil {
			return err
		}

		if reg.MatchString(k) || k == "mirror_link" {
			c.MirrorLinks = append(c.MirrorLinks, v)
			continue
		}
		c.Config[k] = v

	}

	if len(c.Config) == 0 {
		return apperror.ErrReadingConfig
	}

	return nil
}
