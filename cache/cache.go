package cache

type Cache interface {
	Get(key string) interface{}
	Put(key string, val interface{})
	Size() int
}
