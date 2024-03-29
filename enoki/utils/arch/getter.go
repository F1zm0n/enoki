package arch

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/F1zm0n/enoki/enoki/utils/entity"
	apperror "github.com/F1zm0n/enoki/enoki/utils/pkg/AppErro"
	scrapper "github.com/F1zm0n/enoki/enoki/utils/pkg/getter"
)

// func GetPkgList(deps []string) ([]string, error) {
// }

func GetPacman(repo, pkgName, OsArch string) ([]byte, error) {
	pacURl := "https://archlinux.org/packages/" + repo + "/" + OsArch + "/" + pkgName + "/download"

	fmt.Println(pacURl)
	data, err := scrapper.GetFromRepo(pacURl)
	if err != nil {
		if errors.Is(err, apperror.ErrNoSuchPkg) {
			return nil, ErrNoSuchPackage
		}
		return nil, err
	}

	return data, nil
}

// func GetPkgInfoS(ctx context.Context, deps []string) ([]string, error) {
// 	out := make([]string, 0)
//
// 	for _, dep := range deps {
// 		out = append(out, GetPkgInfoPac())
// 	}
//
// 	return out
// }

func GetFromPacmanBoth(repos []string, arch string, pkgName string) ([]byte, error) {
	for _, repo := range repos {
		var dat []byte
		var err error

		if arch == "x86_64" {
			dat, err = GetPacman(repo, pkgName, arch)
			if err != nil {
				dat, err = GetPacman(repo, pkgName, "x86_64")
			}
		} else {
			dat, err = GetPacman(repo, pkgName, arch)
		}

		if err == nil {
			return dat, nil
		}

		if errors.Is(err, ErrNoSuchPackage) {
			// Продолжаем итерации, если ошибка - ErrNoSuchPackage
			continue
		}

		// Возвращаем другие ошибки немедленно
		return nil, err
	}

	return nil, ErrNoSuchPackage
}

func GetPkgInfoPac(repo, OsArch, pkgName string) (entity.ArchInfo, error) {
	url := "https://archlinux.org/packages/" + repo + "/" + OsArch + "/" + pkgName + "/json"

	fmt.Println(url)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return entity.ArchInfo{}, err
	}
	if res.StatusCode != 200 {
		return entity.ArchInfo{}, ErrNoSuchPackage
	}

	var info entity.ArchInfo

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return entity.ArchInfo{}, err
	}

	defer res.Body.Close()

	err = json.Unmarshal(body, &info)
	if err != nil {
		return entity.ArchInfo{}, err
	}
	fmt.Println(info)

	return info, nil
}

func GetPkgInfoBoth(repos []string, arch string, pkgName string) (entity.ArchInfo, error) {
	for _, repo := range repos {
		var dat entity.ArchInfo
		var err error

		if arch == "x86_64" {
			dat, err = GetPkgInfoPac(repo, "x86_64", pkgName)
			if err != nil {
				dat, err = GetPkgInfoPac(repo, "any", pkgName)
			}
		} else {
			dat, err = GetPkgInfoPac(repo, arch, pkgName)
		}

		if err == nil {
			return dat, nil
		}

		if errors.Is(err, ErrNoSuchPackage) {
			// Продолжаем итерации, если ошибка - ErrNoSuchPkg
			continue
		}

		// Возвращаем другие ошибки немедленно
		return entity.ArchInfo{}, err
	}

	return entity.ArchInfo{}, ErrNoSuchPackage
}
