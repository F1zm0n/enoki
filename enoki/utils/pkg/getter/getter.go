package getter

import (
	"io"
	"net/http"

	apperror "github.com/F1zm0n/enoki/enoki/utils/pkg/AppError"
)

// GetReq makes a get request to given link
// and return a slice of bytes, if code is not 200
// it will return apperror.ErrNoSuchPkg
func GetReq(url string) ([]byte, error) {
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
