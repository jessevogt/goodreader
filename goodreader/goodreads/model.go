package goodreader

import (
	"encoding/xml"
	"fmt"
)

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	XMLName xml.Name `xml:"channel"`
	Books   []Book   `xml:"item"`
}

type Book struct {
	XMLName xml.Name `xml:"item"`
	BookId  string   `xml:"book_id"`
	Title   string   `xml:"title"`
	Author  string   `xml:"author_name"`
	ISBN    string   `xml:"isbn"`
}

func ExtractBooks(html []byte) []Book {
	rss := Rss{}
	xml.Unmarshal(html, &rss)
	for _, book := range rss.Channel.Books {
		fmt.Println(book.Title, book.BookId)
	}
	return rss.Channel.Books
}
