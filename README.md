## go-util

golang 实现的工具方法

### cache 缓存
```go
type Cache interface {
	Get(key string) interface{}
	Put(key string, val interface{})
	Size() int
}
```
**lruCache** 最近最少更新



### order 排序

|   排序方法   |  func    |      |
| ---- | ---- | ---- | 
| 冒泡排序   |  Bubble    | |

### stack 栈

### strings 字符串操作
