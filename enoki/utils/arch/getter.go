package arch

import (
	"encoding/json"
	"io"
	"net/http"
	"runtime"

	"github.com/F1zm0n/enoki/enoki/utils/entity"
	scrapper "github.com/F1zm0n/enoki/enoki/utils/pkg/getter"
)

func GetPacman(pkgName, OsArch string) ([]byte, error) {
	pacURl := "https://archlinux.org/packages/extra/" + OsArch + "/" + pkgName + "/download"

	data, err := scrapper.GetFromRepo(pacURl)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetFromPacmanBoth gets the data from pacman either for x86_64 arch or for any based on given
// which variable should be x86_64 for x86_64 arch
func GetFromPacmanBoth(arch string, pkgName string) ([]byte, error) {
	if arch == "x86_64" {
		dat, err := GetPacman(pkgName, "x86_64")
		if err != nil {
			return nil, err
		}
		return dat, nil
	}
	dat, err := GetPacman(pkgName, "any")
	if err != nil {
		return nil, err
	}

	return dat, nil
}

// GetPkgInfoPac gets the json info about package you give and returns 2 channels
// 1 for architercture to download and 2 for info
func GetPkgInfoPac(OsArch, pkgName string) (entity.ArchInfo, error) {
	url := "https://archlinux.org/packages/extra/" + OsArch + "/" + pkgName + "/json"

	req, err := http.NewRequest(http.MethodGet, url, nil)

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return entity.ArchInfo{}, err
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

	return info, nil
}

func GetPkgInfoBoth(pkgName string) (entity.ArchInfo, error) {
	arch := runtime.GOARCH

	if arch == "amd64" {
		dat, err := GetPkgInfoPac("x86_64", pkgName)
		if err != nil {
			return entity.ArchInfo{}, err
		}

		return dat, nil
	}

	dat, err := GetPkgInfoPac("any", pkgName)
	if err != nil {
		return entity.ArchInfo{}, err
	}

	return dat, nil
}
