package pacutils

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/F1zm0n/enoki/enoki/utils/entity"
	apperror "github.com/F1zm0n/enoki/enoki/utils/pkg/AppErro"
	scrapper "github.com/F1zm0n/enoki/enoki/utils/pkg/getter"
)

// GetPkg gets the info about package by this link:
// https://archlinux.org/packages/search/json/?name=%s
func GetPkg(pkgName string) (entity.ArchInfo, error) {
	url := fmt.Sprintf("https://archlinux.org/packages/search/json/?name=%s", pkgName)
	body, err := scrapper.GetReq(url)
	if err != nil {
		return entity.ArchInfo{}, err
	}

	var info entity.PkgPacman

	err = json.Unmarshal(body, &info)
	if err != nil {
		return entity.ArchInfo{}, err
	}

	if len(info.Result) == 0 {
		return entity.ArchInfo{}, apperror.BreakErr
	}
	fmt.Println("res")

	return info.Result[0], nil
}

// GetPkgSo gets the info about package with .so= ending
// by this link: https://archlinux.org/packages/search/json/?name=%s
func GetPkgSo(pkgSo string) (entity.ArchInfo, error) {
	if !strings.Contains(pkgSo, ".so=") {
		return entity.ArchInfo{}, apperror.ErrNoSuchPkg
	}
	url := "https://archlinux.org/packages/search/json/?q=" + pkgSo
	body, err := scrapper.GetReq(url)
	if err != nil {
		return entity.ArchInfo{}, err
	}

	var info entity.PkgPacman

	err = json.Unmarshal(body, &info)
	if err != nil {
		return entity.ArchInfo{}, err
	}

	if len(info.Result) == 0 {
		return entity.ArchInfo{}, apperror.BreakErr
	}

	return info.Result[0], nil
}
