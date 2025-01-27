package engine

import "sync"

type Engine struct {
	mu   sync.RWMutex
	data map[string]string
}

func (e *Engine) Set(key string, value string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.data[key] = value
}

func (e *Engine) Get(key string) (string, bool) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	d, ok := e.data[key]
	return d, ok
}

func (e *Engine) Del(key string) bool {
	e.mu.Lock()
	defer e.mu.Unlock()
	_, ok := e.data[key]
	if ok {
		delete(e.data, key)
		return true
	}
	return false

}

func New() *Engine {
	return &Engine{
		data: make(map[string]string),
	}
}
