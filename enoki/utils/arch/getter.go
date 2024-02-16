package arch

import (
	"errors"
	"fmt"

	"github.com/F1zm0n/enoki/enoki/utils/entity"
	apperror "github.com/F1zm0n/enoki/enoki/utils/pkg/AppErro"
	scrapper "github.com/F1zm0n/enoki/enoki/utils/pkg/getter"
)

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
	return nil, ErrNoSuchPackage
}
