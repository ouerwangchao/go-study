package interview

import (
	"fmt"
	"testing"
)

func TestLru(t *testing.T) {
	var capacity = 3 //容量
	Lc := initLru(capacity)
	fmt.Printf("%v\n", Lc)
	res := Lc.get(2)
	fmt.Printf("%v\n", res)
}

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key   int
	value int
	prev  *DLinkedNode
	next  *DLinkedNode
}

func initLru(c int) LRUCache {
	m := make(map[int]*DLinkedNode, c)
	return LRUCache{
		size:     0,
		capacity: c,
		cache:    m,
		head:     nil,
		tail:     nil,
	}
}

func (lrc *LRUCache) get(key int) int {
	node, ok := lrc.cache[key]
	if !ok {
		return -1
	}
	fmt.Printf("%v存在\n", key)
	//将该节点赋值给头结点
	lrc.setHeader(node)
	return node.value
}

func (lrc *LRUCache) setHeader(node *DLinkedNode) {

}

func (lrc *LRUCache) remove(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
