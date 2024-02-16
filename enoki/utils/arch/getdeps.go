package arch

import (
	"errors"
	"slices"

	pacutils "github.com/F1zm0n/enoki/enoki/utils/arch/utils"
	apperror "github.com/F1zm0n/enoki/enoki/utils/pkg/AppErro"
	utilities "github.com/F1zm0n/enoki/enoki/utils/pkg/utils"
)

// SearchForPkgs searches for all the of the dependencies of dependencies
// and so on and append to a PacmanApp struct slice of string with dependencies
func (a *PacmanApp) SearchForPkgs() error {
	deps := make(map[string]bool)
	deps[a.RootPkg] = true

	for {
		newDeps := make(map[string]bool)

		for dep := range deps {
			pkg, err := pacutils.GetPkg(dep)
			if err != nil {
				if errors.Is(err, apperror.BreakErr) {
					continue
				}
				return err
			}

			for _, name := range pkg.Depends {
				name = pacutils.TrimVersionPacman(name)
				newDeps[name] = true
			}

		}

		if eq := utilities.MapContains(deps, newDeps); len(newDeps) == 0 || eq {
			break
		}
		for name := range newDeps {
			deps[name] = true
		}

	}
	arr := make([]string, 0)
	for key := range deps {
		info, err := pacutils.GetPkgSo(key)
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
