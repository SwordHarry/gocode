package main

import (
	"fmt"
	"math/rand"
	"sort"
)

/**
线性时间选择
给定线性序集中n个元素和一个整数k，1≤k≤n，要求找出这n个元素中第k小的元素
*/

// limit 表示将 多少个 元素分成一组
const limit = 9

// Array 类型，实现 sort 排序接口便于排序
type Array []int

func (a Array) Len() int {
	return len(a)
}

func (a Array) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a Array) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func mid(a Array, i int) int {
	ml := a[i*limit : i*limit+limit]
	sort.Sort(ml)
	return ml[len(ml)/2]
}

// LinearSelect n 为元素个数，k 为第k小元素
func LinearSelect(a Array, n, k int) int {
	var i, j, s int
	// 规模足够小则直接排序
	if n < 243 {
		sort.Sort(a)
		return a[k-1]
	}

	var p, q, r, ml Array
	// 将每组中位数存入数组 ml 中
	for i = 0; i < n/limit; i++ {
		ml = append(ml, mid(a, i))
	}

	// 从 ml 中取中位数的中位数
	// i/2 + i%2 是为了取中位数的中值
	m := LinearSelect(ml, i, i/2+i%2)
	i = 0
	// 根据 m 把原数组划分为 p q r 三份
	for t := 0; t < n; t++ {
		// 比 m 小的放入 p
		if a[t] < m {
			p = append(p, a[t])
			i++
		} else if a[t] == m {
			// 等于 m 的放入 q
			q = append(q, a[t])
			j++
		} else {
			// 小于 m 的放入 r
			r = append(r, a[t])
			s++
		}
	}

	// 第 k 小元素在数组 p 中，继续在 p 中寻找
	if i > k {
		return LinearSelect(p, i, k)
	} else if i+j >= k {
		// m 就是第 k 小元素
		return m
	} else {
		// 第 k 小元素在数组 r 中，继续在 r 中寻找
		return LinearSelect(r, s, k-i-j)
	}
}

func main() {
	a := Array{}
	index := 25
	high := 400
	for i := 0; i < high; i++ {
		a = append(a, i)
	}
	shuffle(a)
	fmt.Printf("总共 %d 个元素，每个组的个数为 %d\n", high, limit)
	for i, v := range a {
		fmt.Printf("a[%d]: %d ", i, v)
	}
	fmt.Println()
	fmt.Printf("第%d小的数字是：%d", index, LinearSelect(a, high, index))
}

// 将原数组打乱
func shuffle(a Array) {
	length := len(a)
	for i := length - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		a[i], a[num] = a[num], a[i]
	}
}
