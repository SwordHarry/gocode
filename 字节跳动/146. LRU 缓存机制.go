package main

/**
运用你所掌握的数据结构，设计和实现一个 LRU (最近最少使用) 缓存机制 。
实现 LRUCache 类：
LRUCache(int capacity) 以正整数作为容量capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。
	当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。

进阶：你是否可以在O(1) 时间复杂度内完成这两种操作？

示例：
输入
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出
[null, null, null, 1, null, -1, null, -1, 3, 4]

解释
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4

提示：
1 <= capacity <= 3000
0 <= key <= 3000
0 <= value <= 104
最多调用 3 * 104 次 get 和 put
*/

type LRUCache struct {
	c          int
	m          map[int]*entry
	start, end *entry
}

type entry struct {
	pre, next *entry
	key, val  int
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		c: capacity,
		m: map[int]*entry{},
	}

	l.start = new(entry)
	l.end = new(entry)
	l.start.next = l.end
	l.end.pre = l.start
	return l
}

func (this *LRUCache) Get(key int) int {
	e, ok := this.m[key]
	if !ok {
		return -1
	}
	this.moveToHead(e)
	return e.val
}

func (this *LRUCache) Put(key int, value int) {
	e, ok := this.m[key]
	// 更新
	if ok {
		e.val = value
		this.moveToHead(e)
		return
	}

	// 插入
	e = &entry{
		key: key,
		val: value,
	}
	// 插入到开头
	sNext := this.start.next
	this.start.next = e
	e.next = sNext
	sNext.pre = e
	e.pre = this.start
	// 更新 map
	this.m[key] = e
	// 淘汰机制
	if len(this.m) > this.c {
		// 队尾淘汰
		e = this.end.pre
		pre := e.pre
		pre.next = this.end
		this.end.pre = pre
		delete(this.m, e.key)
		e = nil
	}
}

func (this *LRUCache) moveToHead(e *entry) {
	// 将找到的 e 移动到队头
	pre := e.pre
	next := e.next
	pre.next = next
	next.pre = pre

	sNext := this.start.next
	this.start.next = e
	e.next = sNext
	e.pre = this.start
	sNext.pre = e
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	var lRUCache = Constructor(2)
	lRUCache.Put(1, 1) // 缓存是 {1=1}
	lRUCache.Put(2, 2) // 缓存是 {1=1, 2=2}
	lRUCache.Get(1)    // 返回 1
	lRUCache.Put(3, 3) // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	lRUCache.Get(2)    // 返回 -1 (未找到)
	lRUCache.Put(4, 4) // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	lRUCache.Get(1)    // 返回 -1 (未找到)
	lRUCache.Get(3)    // 返回 3
	lRUCache.Get(4)    // 返回 4
}
