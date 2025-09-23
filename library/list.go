package library

import (
	"fmt"
	"libraryes/database"
	"libraryes/struct"
	"sync"
)

type Library struct {
	books    map[string]str.Book
	mtx      sync.RWMutex
	postgres database.Postgres
}

func NewLibrary(pg database.Postgres) *Library {
	tempBooks, err := pg.DBExportBooks()
	if err != nil {
		fmt.Println(err)
	}
	return &Library{
		books:    tempBooks,
		postgres: pg,
	}
}

func (l *Library) AddBook(book str.Book) error {
	l.mtx.Lock()
	if _, ok := l.books[book.Title]; ok {
		l.mtx.Unlock()
		return ErrBookAlreadyExists
	}
	l.mtx.Unlock()

	l.postgres.DBInsertBooks(book.Title, book.Author, book.Pages)

	l.books[book.Title] = book

	return nil
}

func (l *Library) GetBook(title string) (str.Book, error) {	
	l.mtx.RLock()
	defer l.mtx.RUnlock()


	book, ok := l.books[title]
	if !ok {
		return str.Book{}, ErrBookNotFound
	}
	return book, nil
}

func (l *Library) ListBooks() map[string]str.Book {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	return l.books
}

func (l *Library) ListUnReadedBooks() map[string]str.Book {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	UnReadedBooks := make(map[string]str.Book)

	for title, book := range l.books {
		if !book.Readed {
			UnReadedBooks[title] = book
		}
	}
	return UnReadedBooks
}

func (l *Library) ListReadedBooks() map[string]str.Book {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	ReadedBooks := make(map[string]str.Book)

	for title, book := range l.books {
		if book.Readed {
			ReadedBooks[title] = book
		}
	}
	return ReadedBooks
}

func (l *Library) ReadBook(title string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	book := l.books[title]

	ReadBook(&book)
	l.postgres.DBReadBook(title)
	l.books[title] = book

	return nil
}

func (l *Library) UnReadBook(title string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	book, ok := l.books[title]
	if !ok {
		return ErrBookNotFound
	}

	UnReadBook(&book)

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

func (l *Library) BoolReadBook(title string) bool {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	book, ok := l.books[title]
	if !ok {
		return false
	}

	b := BoolReadBooks(&book)

	return b
}

func (l *Library) ListBooksAuthor(author string) map[string]str.Book {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	ListBooksAuthor := make(map[string]str.Book)
	for title, book := range l.books {
		if author == book.Author {
			ListBooksAuthor[title] = book
		}
	}
	return ListBooksAuthor
}
