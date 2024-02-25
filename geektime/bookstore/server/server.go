package server

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"hello-go/geektime/bookstore/server/middleware"
	"hello-go/geektime/bookstore/store"
	"net/http"
	"time"
)

type BookStoreServer struct {
	s   store.Store
	srv *http.Server
}

func (booStoreServer *BookStoreServer) createBookHandle(writer http.ResponseWriter, request *http.Request) {
	dec := json.NewDecoder(request.Body)
	var book store.Book
	if err := dec.Decode(&book); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if err := booStoreServer.s.Create(&book); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
}

func (booStoreServer *BookStoreServer) updateBookHandler(writer http.ResponseWriter, request *http.Request) {
	id, ok := mux.Vars(request)["id"]
	if !ok {
		http.Error(writer, "no id found in request", http.StatusBadRequest)
		return
	}

	dec := json.NewDecoder(request.Body)
	var book store.Book
	if err := dec.Decode(&book); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	book.Id = id
	if err := booStoreServer.s.Update(&book); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
}

func (booStoreServer *BookStoreServer) getBookHandler(writer http.ResponseWriter, request *http.Request) {
	id, ok := mux.Vars(request)["id"]
	if !ok {
		http.Error(writer, "no id found in request", http.StatusBadRequest)
		return
	}

	book, err := booStoreServer.s.Get(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	response(writer, book)
}

func (booStoreServer *BookStoreServer) getAllBookHandler(writer http.ResponseWriter, request *http.Request) {
	books, err := booStoreServer.s.GetAll()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	response(writer, books)
}

func (booStoreServer *BookStoreServer) delBookHandler(writer http.ResponseWriter, request *http.Request) {
	id, ok := mux.Vars(request)["id"]
	if !ok {
		http.Error(writer, "no id found in request", http.StatusBadRequest)
		return
	}

	err := booStoreServer.s.Delete(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

}

func response(w http.ResponseWriter, v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func NewBookStoreServer(addr string, s store.Store) *BookStoreServer {
	srv := &BookStoreServer{
		s: s,
		srv: &http.Server{
			Addr: addr,
		},
	}

	router := mux.NewRouter()
	router.HandleFunc("/book", srv.createBookHandle).Methods("POST")
	router.HandleFunc("/book/{id}", srv.updateBookHandler).Methods("POST")
	router.HandleFunc("/book/{id}", srv.getBookHandler).Methods("GET")
	router.HandleFunc("/book", srv.getAllBookHandler).Methods("GET")
	router.HandleFunc("/book/{id}", srv.delBookHandler).Methods("DELETE")

	srv.srv.Handler = middleware.Logging(middleware.Validating(router))
	return srv
}

func (booStoreServer *BookStoreServer) ListenAndServe() (<-chan error, error) {
	var err error
	errChan := make(chan error)
	go func() {
		err = booStoreServer.srv.ListenAndServe()
		errChan <- err
	}()

	select {
	case err = <-errChan:
		return nil, err
	case <-time.After(time.Second):
		return errChan, nil
	}
}

func (bookStoreServer *BookStoreServer) Shutdown(ctx context.Context) error {
	return bookStoreServer.srv.Shutdown(ctx)
}
