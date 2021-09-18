package cache

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLRUCache_Get(t *testing.T) {
	c := 10
	cache := NewLRUCache(c)

	for i := 0; i < c; i++ {
		cache.Put(fmt.Sprintf("key:%d", i), i)
		getVal := cache.Get(fmt.Sprintf("key:%d", i))
		if !reflect.DeepEqual(getVal, i) {
			t.Errorf("TestLRUCache_Get error")
		}
	}
}

func TestLRUCache_Get2(t *testing.T) {
	c := 2
	cache := NewLRUCache(c)

	cache.Put("key:1", 1)
	cache.Put("key:2", 2)
	cache.Put("key:3", 3) // 2, 3
	get1 := cache.Get("key:1")
	if get1 != nil {
		t.Errorf("TestLRUCache_Get2 error")
	}
	cache.Get("key:2")
	cache.Put("key:4", 4) // 2, 4
	cache.Print()
}
