package factory

import (
	"fmt"
	"hello-go/geektime/bookstore/store"
	"sync"
)

var (
	providerMu sync.RWMutex
	providers  = make(map[string]store.Store)
)

func Register(name string, p store.Store) {
	providerMu.Lock()
	defer providerMu.Unlock()
	if p == nil {
		panic("store: Register provider is nil")
	}

	if _, dup := providers[name]; dup {
		panic("store: Register called twice for provider " + name)
	}
	providers[name] = p
}

func New(name string) (store.Store, error) {
	providerMu.RLock()
	p, ok := providers[name]
	providerMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("store: unknown provider %s", name)
	}
	return p, nil
}
