package arch

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/F1zm0n/enoki/enoki/utils/entity"
	apperror "github.com/F1zm0n/enoki/enoki/utils/pkg/AppErro"
)

func GetPkg(pkgName string) (entity.ArchInfo, error) {
	url := "https://archlinux.org/packages/search/json/?name=" + pkgName
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
	for pkg := range deps {
		arr = append(arr, pkg)
	}

	return arr, nil
}
