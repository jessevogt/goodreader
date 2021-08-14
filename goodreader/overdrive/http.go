package goodreader

import (
	"io"
	"net/http"
	"net/url"
)

func GetSearchPage(searchTerm string) ([]byte, error) {
	resp, err := http.Get(
		"https://wplc.overdrive.com/wplc-mcfls/content/search?query=" + url.QueryEscape(searchTerm),
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
