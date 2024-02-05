package scrapper

import (
	"context"
	"io"
	"net/http"

	apperror "github.com/F1zm0n/enoki/enoki/utils/pkg/AppErro"
)

func GetFromRepo(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, apperror.ErrNoSuchPkg
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return body, nil
}
