package main

import (
	"fmt"

	overdrive "github.com/jessevogt/goodreader/goodreader/overdrive"
)

func main() {
	// html, _ := goodreader.GetShelfPage("21054289-jesse-vogt", "to-read", 1)
	// goodreader.ExtractBooks(html)a

	search_html, _ := overdrive.GetSearchPage("Devolution: A Firsthand Account of the Rainier Sasquatch Massacre")
	book_search_results, _ := overdrive.ExtractSearchResults(search_html)
	for i, res := range *book_search_results {
		fmt.Println(i, res)
	}
}
