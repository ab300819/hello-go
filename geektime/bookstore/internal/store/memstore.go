package store

import (
	mystore "hello-go/geektime/bookstore/store"
	factory "hello-go/geektime/bookstore/store/factory"
	"sync"
)

type MemStore struct {
	sync.RWMutex
	books map[string]*mystore.Book
}

func (m *MemStore) Create(book *mystore.Book) error {
	m.RLock()
	defer m.RUnlock()

	if _, ok := m.books[book.Id]; ok {
		return mystore.ErrExist
	}

	nBook := *book
	m.books[book.Id] = &nBook
	return nil
}

func (m *MemStore) Update(book *mystore.Book) error {
	m.RLock()
	defer m.RUnlock()

	oldBook, ok := m.books[book.Id]
	if !ok {
		return mystore.ErrNotFound
	}

	nBook := *oldBook
	if book.Name != "" {
		nBook.Name = book.Name
	}

	if book.Press != "" {
		nBook.Press = book.Press
	}

	m.books[book.Id] = &nBook
	return nil

}

func (m *MemStore) Get(s string) (mystore.Book, error) {
	m.RLock()
	defer m.RUnlock()

	t, ok := m.books[s]
	if ok {
		return *t, nil
	}

	return mystore.Book{}, mystore.ErrNotFound
}

func (m *MemStore) GetAll() ([]mystore.Book, error) {
	m.RLock()
	defer m.RUnlock()

	allBooks := make([]mystore.Book, 0, len(m.books))
	for _, book := range m.books {
		allBooks = append(allBooks, *book)
	}
	return allBooks, nil
}

func (m *MemStore) Delete(id string) error {
	m.RLock()
	defer m.RUnlock()

	if _, ok := m.books[id]; !ok {
		return mystore.ErrNotFound
	}

	delete(m.books, id)
	return nil
}

func init() {
	factory.Register("mem", &MemStore{
		books: make(map[string]*mystore.Book),
	})
}
