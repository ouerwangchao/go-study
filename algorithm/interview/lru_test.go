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
	Lc.put(3,2)
	fmt.Printf("%v\n", Lc)
	v := Lc.get(3)
	t.Log(v)
}

type LRUCache struct {
	capacity   int
	cache      map[int]*LinkNode
	head, tail *LinkNode
}

type LinkNode struct {
	key   int
	value int
	pre   *LinkNode
	next  *LinkNode
}

func initLru(capacity int) LRUCache {
	m := make(map[int]*LinkNode, capacity)
	head := &LinkNode{0, 0, nil, nil}
	tail := &LinkNode{0, 0, nil, nil}
	head.next = tail
	tail.pre = head
	return LRUCache{
		capacity: capacity,
		cache:    m,
		head:     head,
		tail:     tail,
	}
}

//添加节点到头结点之后
func (lrc *LRUCache) AddNode(node *LinkNode) {
	node.pre = lrc.head
	node.next = lrc.head.next
	lrc.head.next = node
	node.next.pre = node
}

//移除节点
func (lrc *LRUCache) RemoveNode(node *LinkNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

//移动到头结点之后
func (lrc *LRUCache) MoveToHead(node *LinkNode) {
	lrc.RemoveNode(node)
	lrc.AddNode(node)
}

func (lrc *LRUCache) get(key int) int {
	node, ok := lrc.cache[key]
	if !ok {
		return -1
	}
	fmt.Printf("%v存在\n", key)
	//将该节点赋值给头结点
	lrc.MoveToHead(node)
	return node.value
}

func (lrc *LRUCache) put(key, value int) {
	if node, ok := lrc.cache[key]; ok {
		//存在则更新
		node.value = value
		lrc.MoveToHead(node)
	} else {
		//判断容量
		node = &LinkNode{key, value, nil, nil}
		lrc.cache[key] = node
		lrc.AddNode(node)
		if len(lrc.cache) >= lrc.capacity {
			//删除一个节点
			delete(lrc.cache, lrc.tail.pre.value)
			lrc.RemoveNode(lrc.tail.pre)
		}
	}
}
