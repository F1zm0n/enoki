package arch

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"

	apperror "github.com/F1zm0n/enoki/enoki/utils/pkg/AppError"
	scrapper "github.com/F1zm0n/enoki/enoki/utils/pkg/getter"
)

// GetMirrorsPacman gets the fastest mirrors to respond and appends it to
// PacmanApp.MirrorLinks string slice
func (a *PacmanApp) GetMirrorsPacman() error {
	url := "https://archlinux.org/mirrorlist/?country=all&protocol=http&protocol=https&ip_version=4&use_mirror_status=on"
	dat, err := scrapper.GetReq(url)
	if err != nil {
		return err
	}

	arr, err := ParseMirrors(dat)
	if err != nil {
		return err
	}
	a.MirrorLinks = arr
	return nil
}

// GetMirrorCountryHost makes request to arch mirror link repository and appends to
// PacmanApp struct MirrorLinks field
func (a *PacmanApp) GetMirrorCountryHost(conf map[string]string) error {
	if country, ok := conf["country"]; ok {

		url := fmt.Sprintf(
			"https://archlinux.org/mirrorlist/?country=%s&protocol=https&protocol=http&ip_version=4&ip_version=6",
			country,
		)
		dat, err := scrapper.GetReq(url)
		if err != nil {
			return err
		}

		arr, err := ParseMirrors(dat)
		if err != nil {
			return err
		}
		if index, ok := conf["number"]; ok {
			idx, err := strconv.Atoi(index)
			if err != nil {
				fmt.Println("config error")
				return err
			}
			fmt.Println("get some normal mirror links here arch/conf.go", arr[idx+1])
			return nil
		} else {
			return nil
		}

	}
	if link, ok := conf["link"]; ok {
		arr := strings.Split(link, "/")
		fmt.Println("fix this in arch/conf", arr)
		return nil
	}
	fmt.Println("config error")
	return apperror.ErrReadingConfig
}

// ParseMirrors parses mirror links and return slice of links
// to get data from
func ParseMirrors(data []byte) ([]string, error) {
	f := bytes.NewReader(data)

	arr := make([]string, 0, 700)

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
