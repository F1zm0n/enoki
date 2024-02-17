package arch

import (
	"errors"
	"fmt"

	"github.com/F1zm0n/enoki/enoki/utils/entity"
	apperror "github.com/F1zm0n/enoki/enoki/utils/pkg/AppError"
	scrapper "github.com/F1zm0n/enoki/enoki/utils/pkg/getter"
)

// GetFromPacman gets package from pacman repository by mirror link
// so PacmanApp.MirrorLinks slice is required to make request try and get the package
func (a *PacmanApp) GetFromPacman(info entity.ArchInfo) ([]byte, error) {
	for _, hoster := range a.MirrorLinks {
		for _, archi := range a.Arches {
			pacURL := fmt.Sprintf(
				"%s/archlinux/%s/os/%s/%s",
				hoster,
				info.Repo,
				archi,
				info.Filename,
			)
			body, err := scrapper.GetReq(pacURL)
			if err != nil {
				if errors.Is(err, apperror.ErrNoSuchPkg) {
					continue
				}
				return nil, err
			} else {
				return body, nil
			}
		}
	}
	return nil, apperror.ErrNoSuchPkg
}
