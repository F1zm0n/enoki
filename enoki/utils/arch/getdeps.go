package arch

import (
	"fmt"
	"time"

	"github.com/F1zm0n/enoki/enoki/utils/entity"
)

func InstDepsPacman(
	timeout time.Duration,
	repos []string,
	pkgInfo entity.ArchInfo,
	pkgArch string,
	dir string,
	arr []string,
) ([]string, error) {
	var err error = nil

	for _, pkg := range pkgInfo.Depends {
		arr, err = UnpakAndInstPacman(timeout, arr, repos, pkgArch, pkg, dir)
		fmt.Println("deps")
		if err != nil {
			fmt.Println("deps,error")
			return nil, err
		}
	}
	return arr, nil
}
