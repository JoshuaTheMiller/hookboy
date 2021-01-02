package util

import (
	"io/ioutil"
	"net/http"
)

// HTTPGetter retrieves a document from a provided URL as a string
type HTTPGetter interface {
	Get(url string) (string, error)
}

type hTTPGetter struct {
}

// DefaultHTTPGetter returns a default HTTPGetter implementation
func DefaultHTTPGetter() HTTPGetter {
	return hTTPGetter{}
}

func (h hTTPGetter) Get(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
