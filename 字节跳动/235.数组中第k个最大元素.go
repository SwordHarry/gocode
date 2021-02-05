package main

/**
在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

示例 1:
输入: [3,2,1,5,6,4] 和 k = 2
输出: 5

示例 2:
输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4

说明:
你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
*/

// 解法1：采用 堆，优先队列 即可实现
import "container/heap"

type myHeap struct {
	queue []int
}

func (m *myHeap) Len() int {
	return len(m.queue)
}

func (m *myHeap) Less(i, j int) bool {
	return m.queue[i] > m.queue[j]
}

func (m *myHeap) Swap(i, j int) {
	m.queue[i], m.queue[j] = m.queue[j], m.queue[i]
}

func (m *myHeap) Push(x interface{}) {
	m.queue = append(m.queue, x.(int))
}

func (m *myHeap) Pop() interface{} {
	last := len(m.queue) - 1
	result := m.queue[last]
	m.queue = m.queue[:last]
	return result
}

func findKthLargest(nums []int, k int) int {
	mh := &myHeap{queue: nums}
	heap.Init(mh)
	var result interface{}
	for i := 0; i < k; i++ {
		result = heap.Pop(mh)
	}
	return result.(int)
}

// 解法2：采用 快排分治 的思想
func findKthLargest2(nums []int, k int) int {
	var big []int
	var small []int
	mid := nums[0]
	nums = nums[1:]
	for _, v := range nums {
		if v >= mid {
			big = append(big, v)
		} else {
			small = append(small, v)
		}
	}
	bigLen := len(big)
	if bigLen > k-1 {
		return findKthLargest2(big, k)
	} else if bigLen == k-1 {
		return mid
	} else {
		// 除掉 mid 和 big
		return findKthLargest2(small, k-bigLen-1)
	}
}

func main() {

}
