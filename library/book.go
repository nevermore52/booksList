package library

import (
	"libraryes/struct"
	"time"
)
func NewBook(title string, author string, pages int) str.Book {
	return str.Book{
		Title: title,
		Author: author,
		Pages: pages,
		Readed: false,

		Timeadd: time.Now(),
		Timereaded: nil,
	}
}

func ReadBook() {
	TimeRead := time.Now()
	b.Readed = true
	b.Timereaded = &TimeRead
}

func (b *str.Book) UnReadBook() {
	b.Readed = false
	b.Timereaded = nil
}

func (b *Book) BoolReadBooks() bool {
	return b.Readeed
}