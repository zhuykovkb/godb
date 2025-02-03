package inMemory

type Engine struct {
	data map[string]string
}

func (e *Engine) Set(key string, value string) {
	e.data[key] = value
}

func (e *Engine) Get(key string) (string, bool) {
	d, ok := e.data[key]
	return d, ok
}

func (e *Engine) Del(key string) {
	delete(e.data, key)
}

func NewEngine() *Engine {
	return &Engine{
		data: make(map[string]string),
	}
}
