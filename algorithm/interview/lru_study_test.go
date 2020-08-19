/**
  @author：wangchao
  @date：2020/8/18
  @note：LRU淘汰算法
**/
package interview

import (
	"testing"
)

/**
分析：LRU淘汰算法：淘汰最近最少使用数据
维护一个循环链表和map即可实现快速增删改查操作
*/

//节点
type LinkNode struct {
	Key   int
	Value int
	Prev  *LinkNode
	Next  *LinkNode
}

//LRU结构
type LruCache struct {
	Cap      int
	CacheMap map[int]*LinkNode
	Head     *LinkNode
	Tail     *LinkNode
}

var l LruCache

//初始化LRU
func InitLru(c int) {
	head := &LinkNode{0, 0, nil, nil}
	tail := &LinkNode{0, 0, nil, nil}
	head.Next = tail
	tail.Prev = head
	cacheMap := make(map[int]*LinkNode, c)
	l = LruCache{c, cacheMap, head, tail}
}

//将节点插入循环链表头结点之后
func (lrc *LruCache) addNodeFirst(node *LinkNode) {
	node.Prev = lrc.Head
	node.Next = lrc.Head.Next
	lrc.Head.Next = node
	node.Next.Prev = node
}

//删除节点
func (lrc *LruCache) RemoveNode(node *LinkNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

//将节点移动到循环链表头结点之后
func (lrc *LruCache) MoveToFirst(node *LinkNode) {
	lrc.RemoveNode(node)
	lrc.addNodeFirst(node)
}

func (lrc *LruCache) get(key int) int {
	if node, ok := lrc.CacheMap[key]; ok {
		lrc.MoveToFirst(node)
		return node.Value
	}
	return -1
}

func (lrc *LruCache) put(key, value int) {
	//判断当前链表是否包含这个key
	if node, ok := lrc.CacheMap[key]; ok {
		//包含 更新值 将节点移动到头部
		node.Value = value
		lrc.CacheMap[key] = node
		lrc.MoveToFirst(node)
	} else {
		//不包含 判断容量是否足够
		if len(lrc.CacheMap) >= lrc.Cap {
			//容量不足，删除最后一个节点
			node = lrc.Tail.Prev
			lrc.RemoveNode(node)
			delete(lrc.CacheMap, node.Key)
		}
		node = &LinkNode{key, value, nil, nil}
		lrc.addNodeFirst(node)
		lrc.CacheMap[key] = node
	}
}

func TestLruStudy(t *testing.T) {
	InitLru(4)
	t.Log(l)
	a := l.get(5)
	t.Log(a)
	l.put(5, 3)
	t.Log(l)
	t.Log(l.get(5))
}
