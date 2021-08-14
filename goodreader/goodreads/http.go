package goodreader

import (
	"fmt"
	"io"
	"net/http"
)

func GetShelfPage(account string, shelf string, page int) ([]byte, error) {
	booksPerPage := 100
	resp, err := http.Get(
		fmt.Sprintf(
			"https://www.goodreads.com/review/list_rss/%s?utf8=✓&order=d&shelf=%s&sort=date_added&per_page=%d&page=%d",
			account, shelf, booksPerPage, page,
		),
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
