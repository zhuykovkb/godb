package inMemory

import "sync"

type Engine struct {
	m    sync.RWMutex
	data map[string]string
}

func (e *Engine) Set(key string, value string) {
	e.m.Lock()
	defer e.m.Unlock()
	e.data[key] = value
}

func (e *Engine) Get(key string) (string, bool) {
	e.m.Lock()
	defer e.m.Unlock()
	d, ok := e.data[key]
	return d, ok
}

func (e *Engine) Del(key string) {
	e.m.Lock()
	defer e.m.Unlock()
	delete(e.data, key)
}

func NewEngine() *Engine {
	return &Engine{
		data: make(map[string]string),
	}
}
