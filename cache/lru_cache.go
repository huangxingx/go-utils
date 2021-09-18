package cache

import "fmt"

// implement Cache interface
var _ Cache = &lruCache{}

type lruCache struct {
	Cap    int //容量
	Bucket map[string]*Node
	Head   *Node //头节点
	Tail   *Node //尾节点
}

type Node struct {
	Key  string
	Val  interface{}
	Pre  *Node
	Next *Node
}

// NewLRUCache new lruCache by cap
func NewLRUCache(cap int) *lruCache {
	bucket := make(map[string]*Node, cap)
	return &lruCache{
		Cap:    cap,
		Bucket: bucket,
	}
}

// Get by key
func (l *lruCache) Get(key string) interface{} {
	if node, ok := l.Bucket[key]; ok {
		l.refreshNode(node)
		return node.Val
	}
	return nil
}

// Put key value
func (l *lruCache) Put(key string, val interface{}) {
	// exists refresh value
	if node, ok := l.Bucket[key]; ok {
		node.Val = val
		l.refreshNode(node)
		return
	}
	// not exists
	if len(l.Bucket) >= l.Cap {
		l.removeNode(l.Head)
	}
	node := &Node{Key: key, Val: val}
	l.addNode(node)
	return
}

//Print cache value
func (l *lruCache) Print() {
	root := l.Head
	for root != nil {
		fmt.Printf("current Node key is[%+v],value is [%+v]\n", root.Key, root.Val)
		root = root.Next
	}
	return
}

//Size get cache size
func (l *lruCache) Size() int {
	return len(l.Bucket)
}

func (l *lruCache) addNode(node *Node) {
	// tail
	if l.Tail != nil {
		l.Tail.Next = node
		node.Pre = l.Tail
		node.Next = nil
	}
	l.Tail = node
	// head
	if l.Head == nil {
		l.Head = node
		node.Next = nil
	}
	l.Bucket[node.Key] = node
	return
}

//removeNode remove Node
func (l *lruCache) removeNode(node *Node) {
	if l.Tail == node {
		l.Tail = l.Tail.Pre
	} else if l.Head == node {
		l.Head = l.Head.Next
		l.Head.Pre = nil
	} else {
		node.Pre.Next = node.Next
		node.Next.Pre = node.Pre
	}
	delete(l.Bucket, node.Key)
	return
}

//refreshNode move Node to tail
func (l *lruCache) refreshNode(node *Node) {
	if l.Tail == node {
		return
	}
	l.removeNode(node)
	l.addNode(node)
	return
}
