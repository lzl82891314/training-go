package store

import (
	mystore "bookstore/store"
	"bookstore/store/factory"
	"sync"
)

func init() {
	factory.Register("mem", &MemStore{
		books: make(map[string]*mystore.Book),
	})
}

type MemStore struct {
	sync.RWMutex
	books map[string]*mystore.Book
}

func (ms *MemStore) Create(book *mystore.Book) error {
	ms.Lock()
	defer ms.Unlock()

	if _, ok := ms.books[book.Id]; ok {
		return mystore.ErrExist
	}

	newBook := *book
	ms.books[book.Id] = &newBook
	return nil
}

func (ms *MemStore) Update(book *mystore.Book) error {
	ms.Lock()
	defer ms.Unlock()

	preBook, ok := ms.books[book.Id]
	if !ok {
		return mystore.ErrNotFound
	}

	newBook := *preBook
	if book.Name != "" {
		newBook.Name = book.Name
	}
	if book.Authors != nil {
		newBook.Authors = book.Authors
	}
	if book.Press != "" {
		newBook.Press = book.Press
	}
	ms.books[book.Id] = &newBook
	return nil
}

func (ms *MemStore) Get(id string) (mystore.Book, error) {
	ms.Lock()
	defer ms.Unlock()

	t, ok := ms.books[id]
	if !ok {
		return mystore.Book{}, mystore.ErrNotFound
	}
	return *t, nil
}

func (ms *MemStore) Delete(id string) error {
	ms.Lock()
	defer ms.Unlock()

	if _, ok := ms.books[id]; !ok {
		return mystore.ErrExist
	}
	delete(ms.books, id)
	return nil
}

func (ms *MemStore) GetAll() ([]mystore.Book, error) {
	ms.Lock()
	defer ms.Unlock()

	books := make([]mystore.Book, 0, len(ms.books))
	for _, book := range ms.books {
		books = append(books, *book)
	}
	return books, nil
}
