package arch

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/F1zm0n/enoki/enoki/utils/entity"
	scrapper "github.com/F1zm0n/enoki/enoki/utils/pkg/getter"
)

func GetPacman(ctx context.Context, pkgName, OsArch string) ([]byte, error) {
	pacURl := "https://archlinux.org/packages/extra/" + OsArch + "/" + pkgName + "/download"

	data, err := scrapper.GetFromRepo(ctx, pacURl)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetFromPacmanBoth(ctx context.Context, arch string, pkgName string) ([]byte, error) {
	dat, err := GetPacman(ctx, pkgName, arch)
	if err != nil {
		return nil, err
	}

	return dat, nil
}

// GetPkgInfoPac gets the json info about package you give and returns 2 channels
// 1 for architercture to download and 2 for info
func GetPkgInfoPac(ctx context.Context, repo, OsArch, pkgName string) (entity.ArchInfo, error) {
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

func GetPkgInfoBoth(
	timeout time.Duration,
	repos []string,
	arch string,
	pkgName string,
) (entity.ArchInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	for _, repo := range repos {
		if arch == "x86_64" {

			dat, err := GetPkgInfoPac(ctx, repo, arch, pkgName)
			if err != nil {
				dat, err := GetPkgInfoPac(ctx, repo, "any", pkgName)
				if err == nil {
					return dat, nil
				}
			}
			return dat, nil
		}

		dat, err := GetPkgInfoPac(ctx, repo, arch, pkgName)
		if err == nil {
			return dat, nil
		}

	}

	return entity.ArchInfo{}, ErrNoSuchPackage
}
