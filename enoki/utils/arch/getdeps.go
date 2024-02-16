package arch

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strings"

	"github.com/F1zm0n/enoki/enoki/utils/entity"
	apperror "github.com/F1zm0n/enoki/enoki/utils/pkg/AppErro"
)

func GetPkgSo(pkgSo string) (entity.ArchInfo, error) {
	if !strings.Contains(pkgSo, ".so=") {
		return entity.ArchInfo{}, ErrSoPackage
	}
	url := "https://archlinux.org/packages/search/json/?q=" + pkgSo
	req, err := http.NewRequest(http.MethodGet, url, nil)

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return entity.ArchInfo{}, err
	}
	// if res.StatusCode != 200 {
	// 	return ArchInfo{}, ErrNoSuchPackage
	// }

	var info entity.PkgPacman

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return entity.ArchInfo{}, err
	}

	defer res.Body.Close()

	err = json.Unmarshal(body, &info)
	if err != nil {
		return entity.ArchInfo{}, err
	}
	if len(info.Result) == 0 {
		return entity.ArchInfo{}, apperror.BreakErr
	}

	return info.Result[0], nil
}

func GetPkg(pkgName string) (entity.ArchInfo, error) {
	url := fmt.Sprintf("https://archlinux.org/packages/search/json/?name=%s", pkgName)
	req, err := http.NewRequest(http.MethodGet, url, nil)

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return entity.ArchInfo{}, err
	}
	if res.StatusCode != 200 {
		return entity.ArchInfo{}, ErrNoSuchPackage
	}

	var info entity.PkgPacman

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return entity.ArchInfo{}, err
	}

	defer res.Body.Close()

	err = json.Unmarshal(body, &info)
	if err != nil {
		return entity.ArchInfo{}, err
	}
	if len(info.Result) <= 0 {
		return entity.ArchInfo{}, apperror.BreakErr
	}
	fmt.Println("res")

	return info.Result[0], nil
}

func mapContains(m1, m2 map[string]bool) bool {
	for k, v := range m2 {
		if val, ok := m1[k]; !ok || val != v {
			return false
		}
	}
	return true
}

func SearchForPkgs(root string) ([]string, error) {
	deps := make(map[string]bool)
	deps[root] = true

	for {
		newDeps := make(map[string]bool)

		for dep := range deps {
			pkg, err := GetPkg(dep)
			if err != nil {
				if errors.Is(err, apperror.BreakErr) {
					continue
				}
				return nil, err
			}

			for _, name := range pkg.Depends {
				name = TrimVersionPacman(name)
				newDeps[name] = true
			}

		}

		if eq := mapContains(deps, newDeps); len(newDeps) == 0 || eq {
			break
		}
		for name := range newDeps {
			deps[name] = true
		}

	}
	arr := make([]string, 0)
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

	return res, nil
}

func TrimSoPacman(s string) string {
	idx := strings.Index(s, ".so")
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
}

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
