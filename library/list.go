package library

import (
	"libraryes/database"
	"sync"

	"github.com/jmoiron/sqlx"
)

type Library struct {
	books  	map[string]Book
	mtx		sync.RWMutex
	postgres database.Postgres
}
func NewLibrary() *Library{
	return  &Library{
		books: make(map[string]Book),
	}
}

func (l *Library) AddBook(book Book, db *sqlx.DB) error {
	l.mtx.Lock()
	
	if _, ok := l.books[book.Title]; ok{
		l.mtx.Unlock()
		return ErrBookAlreadyExists
	}
	l.mtx.Unlock()		
	
    l.postgres.DBInsertBooks(book.Title,book.Author,book.Pages)

	l.books[book.Title] = book

	

	return nil
}

func (l *Library) GetBook(title string) (Book, error) {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	book , ok  := l.books[title]
	if !ok{
		return Book{}, ErrBookNotFound
	}
	return book, nil
}

func (l *Library) ListBooks() map[string]Book {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	return l.books
}

func (l *Library) ListUnReadedBooks() map[string]Book {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	UnReadedBooks := make(map[string]Book)

	for title, book := range l.books{
		if !book.Readeed{
		UnReadedBooks[title] = book
		}
	}
	return UnReadedBooks
}

func (l *Library) ListReadedBooks() map[string]Book {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	ReadedBooks := make(map[string]Book)

	for title, book := range l.books{
		if book.Readeed{
			ReadedBooks[title] = book
		}
	}
	return ReadedBooks
}

func (l *Library) ReadBook(title string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	book, ok := l.books[title]
	if !ok {
		return ErrBookNotFound
	}

	book.ReadBook()

	l.books[title] = book

	return nil
}

func (l *Library) UnReadBook(title string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	book, ok := l.books[title]
	if !ok{
		return ErrBookNotFound
	}

	book.UnReadBook()

	l.books[title] = book

	return nil
}

func (l *Library) DeleteBook(title string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	_, ok := l.books[title]
	if !ok {
		return ErrBookNotFound
	}

	delete(l.books, title)

	return nil
}

func (l *Library) BoolReadBook(title string) (bool) {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	book, ok := l.books[title]
	if !ok {
		return false
	}

	b := book.BoolReadBooks()

	return b
}

func (l *Library) ListBooksAuthor(author string) map[string]Book {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	ListBooksAuthor := make(map[string]Book)
	for title, book := range  l.books {
		if author == book.Author {
			ListBooksAuthor[title] = book
		}
	}
	return ListBooksAuthor
}