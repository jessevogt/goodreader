package goodreader

import (
	"bytes"
	"encoding/json"
	"regexp"
)

type BookSearchResultJson struct {
	Author      string   `json:"firstCreatorName"`
	Title       string   `json:"sortTitle"`
	ItemType    ItemType `json:"type"`
	IsAvailable bool     `json:"isAvailable"`
}

type BookSearchResult struct {
	Author      string
	Title       string
	ItemType    string
	IsAvailable bool
}

type ItemType struct {
	Id string `json:"id"`
}

type MalformedCollectionError struct{}

func (e *MalformedCollectionError) Error() string {
	return "Could not find or collection payload in html was malformed"
}

func ExtractSearchResults(html []byte) (*[]BookSearchResult, error) {
	regex := regexp.MustCompile(`(?s)window\.OverDrive\.titleCollection = (\[.+)`)
	res := regex.FindSubmatchIndex(html)
	book_search_results_json, err := extract(html[res[2]:])
	if err != nil {
		return nil, err
	}
	var book_search_results []BookSearchResult
	for _, res := range *book_search_results_json {
		book_search_results = append(book_search_results, BookSearchResult{
			Author:      res.Author,
			Title:       res.Title,
			ItemType:    res.ItemType.Id,
			IsAvailable: res.IsAvailable,
		})
	}
	return &book_search_results, nil
}

func extract(html []byte) (*[]BookSearchResultJson, error) {
	closingSquartBracket := []byte("]")

	endPos := bytes.Index(html, closingSquartBracket)
	if endPos == -1 {
		return nil, &MalformedCollectionError{}
	}

	var bookSearchResults []BookSearchResultJson

	for json.Unmarshal(html[0:endPos+1], &bookSearchResults) != nil {
		nextPos := bytes.Index(html[endPos+1:], closingSquartBracket)
		if nextPos == -1 {
			return nil, &MalformedCollectionError{}
		}
		endPos += nextPos + 1
	}

	if endPos == -1 {
		return nil, &MalformedCollectionError{}
	}

	return &bookSearchResults, nil
}
