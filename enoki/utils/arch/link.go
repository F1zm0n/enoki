package arch

import (
	"fmt"
	"strconv"
	"strings"

	pacutils "github.com/F1zm0n/enoki/enoki/utils/arch/utils"
	apperror "github.com/F1zm0n/enoki/enoki/utils/pkg/AppErro"
	scrapper "github.com/F1zm0n/enoki/enoki/utils/pkg/getter"
)

// GetMirrorHost makes request to arch mirror link repository and appends to
// PacmanApp struct MirrorLinks field
func (a *PacmanApp) GetMirrorHost(conf map[string]string) error {
	if country, ok := conf["country"]; ok {

		url := fmt.Sprintf(
			"https://archlinux.org/mirrorlist/?country=%s&protocol=https&protocol=http&ip_version=4&ip_version=6",
			country,
		)
		dat, err := scrapper.GetReq(url)
		if err != nil {
			return err
		}

		arr, err := pacutils.ParseMirrors(dat)
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
	return apperror.BreakErr
}
