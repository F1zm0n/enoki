package arch

import (
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/F1zm0n/enoki/enoki/utils/entity"
	apperror "github.com/F1zm0n/enoki/enoki/utils/pkg/AppError"
	"github.com/F1zm0n/enoki/enoki/utils/pkg/getter"
)

// SearchForPkgs method searches for all the of the dependencies of dependencies
// and so on from pacman repository using 2 function to search for packages:
//
// GetPkgSo and GetPkg functions, and after that appends
// to a PacmanApp struct slice of string with dependencies
//
// returns apperror.ErrNoSuchPkg if no package was found
func (a *PacmanApp) SearchForPkgs() error {
	deps := make(map[string]bool)
	deps[a.RootPkg] = true

	for {
		newDeps := make(map[string]bool)

		for dep := range deps {
			pkg, err := GetPkg(dep)
			if err != nil {
				if errors.Is(err, apperror.ErrNoSuchPkg) {
					continue
				}
				return err
			}

			for _, name := range pkg.Depends {
				name = TrimVersionPacman(name)
				newDeps[name] = true
			}

		}

		if eq := MapContains(deps, newDeps); len(newDeps) == 0 || eq {
			break
		}
		for name := range newDeps {
			deps[name] = true
		}

	}
	arr := make([]string, 0, 200)
	for key := range deps {
		info, err := GetPkgSo(key)
		if err != nil {
			arr = append(arr, key)
		} else {
			arr = append(arr, info.PkgName)
		}
	}
	res := []string{}
	for _, val := range arr {
		if !slices.Contains[[]string](res, val) {
			res = append(res, val)
		}
	}
	a.Deps = res

	return nil
}

// TrimVersionPacman trims all versioned package suffix e.g. linux-api-headers>=4.10
// will be equal to linux-api-headers
func TrimVersionPacman(s string) string {
	idx := strings.Index(s, ">")
	if idx == -1 {
		idx = strings.Index(s, "<")
		if idx == -1 {
			return s
		}
		str := strings.Split(s, "")

		for i := idx; i < len(s); i++ {
			str[i] = " "
		}

		s = strings.Join(str, "")

		s = strings.TrimSpace(s)

		return s

	} else {

		str := strings.Split(s, "")

		for i := idx; i < len(s); i++ {
			str[i] = " "
		}
		s = strings.Join(str, "")
		s = strings.TrimSpace(s)
		return s
	}
}

// GetPkg gets the info about package by this link:
// https://archlinux.org/packages/search/json/?name=%s
func GetPkg(pkgName string) (entity.ArchInfo, error) {
	url := fmt.Sprintf("https://archlinux.org/packages/search/json/?name=%s", pkgName)
	body, err := getter.GetReq(url)
	if err != nil {
		return entity.ArchInfo{}, err
	}

	var info entity.PkgPacman

	err = json.Unmarshal(body, &info)
	if err != nil {
		return entity.ArchInfo{}, err
	}

	if len(info.Result) == 0 {
		return entity.ArchInfo{}, apperror.ErrNoSuchPkg
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
	body, err := getter.GetReq(url)
	if err != nil {
		return entity.ArchInfo{}, err
	}

	var info entity.PkgPacman

	err = json.Unmarshal(body, &info)
	if err != nil {
		return entity.ArchInfo{}, err
	}

	if len(info.Result) == 0 {
		return entity.ArchInfo{}, apperror.ErrNoSuchPkg
	}

	return info.Result[0], nil
}

// MapContains checks if m2 map is in m1 map
func MapContains(m1, m2 map[string]bool) bool {
	for k, v := range m2 {
		if val, ok := m1[k]; !ok || val != v {
			return false
		}
	}
	return true
}
