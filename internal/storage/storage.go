package storage

type Storage interface {
	Get(key string) (string, bool)
	Set(key string, value string)
	Del(key string)
}
