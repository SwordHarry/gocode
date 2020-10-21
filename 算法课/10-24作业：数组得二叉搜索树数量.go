package main

import (
	"fmt"
	"sort"
)

/**
1.给出一个数组，算出其能构成的不同二叉搜索树结构的数量
*/

type array []int

func (a array) Len() int {
	return len(a)
}

func (a array) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a array) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// 二叉树节点
type node struct {
	val         int
	left, right *node
}

func genBST(a array, l, h int) []*node {
	var trees []*node
	// 边界条件1
	if l == h-1 {
		trees = append(trees, &node{
			val: a[l],
		})
		return trees
	}
	// 边界条件2
	if l == h {
		trees = append(trees, nil)
		return trees
	}
	// 递归构造 二叉搜索树
	for i := l; i < h; i++ {
		leftTree := genBST(a, l, i)
		rightTree := genBST(a, i+1, h)
		for _, lt := range leftTree {
			for _, rt := range rightTree {
				root := &node{
					val: a[i],
				}
				root.left = lt
				root.right = rt
				trees = append(trees, root)
			}
		}
	}

	return trees
}

func main() {
	a := array{2, 1, 3, 6, 7, 4, 9}
	// 先对数组进行排序
	sort.Sort(a)
	result := genBST(a, 0, len(a))
	fmt.Println("数组", a, "的二叉搜索树数量为：", len(result))
}
