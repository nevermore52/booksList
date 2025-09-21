package library

import "time"

type Book struct{
	Title		string
	Author		string
	Pages		int
	Readeed		bool
	
	TimeAdd		time.Time
	TimeReaded	*time.Time
}

func NewBook(title string, author string, pages int) Book {
	return Book{
		Title: title,
		Author: author,
		Pages: pages,
		Readeed: false,

		TimeAdd: time.Now(),
		TimeReaded: nil,
	}
}

func (b *Book) ReadBook() {
	TimeRead := time.Now()
	b.Readeed = true
	b.TimeReaded = &TimeRead
}

func (b *Book) UnReadBook() {
	b.Readeed = false
	b.TimeReaded = nil
}

func (b *Book) BoolReadBooks() bool {
	return b.Readeed
}