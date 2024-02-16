package arch

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"

	apperror "github.com/F1zm0n/enoki/enoki/utils/pkg/AppErro"
	scrapper "github.com/F1zm0n/enoki/enoki/utils/pkg/getter"
)

func GetMirrorHost(conf map[string]string) (string, error) {
	if country, ok := conf["country"]; ok {

		url := fmt.Sprintf(
			"https://archlinux.org/mirrorlist/?country=%s&protocol=https&protocol=http&ip_version=4&ip_version=6",
			country,
		)
		dat, err := scrapper.GetReq(url)
		if err != nil {
			return "", err
		}

		arr, err := ParseMirrors(dat)
		if err != nil {
			return "", err
		}
		if index, ok := conf["number"]; ok {
			idx, err := strconv.Atoi(index)
			if err != nil {
				fmt.Println("config error")
				return "", err
			}
			return arr[idx+1], nil
		} else {
			return arr[1], nil
		}

	}
	if link, ok := conf["link"]; ok {
		arr := strings.Split(link, "/")
		return arr[2], nil
	}
	fmt.Println("config error")
	return "", apperror.BreakErr
}

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
