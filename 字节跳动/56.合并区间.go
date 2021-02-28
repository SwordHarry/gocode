package main

import "sort"

/**
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
示例 1：
输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

示例 2：
输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

提示：
1 <= intervals.length <= 104
intervals[i].length == 2
0 <= starti <= endi <= 104
*/

/**
思路：直接模拟，但是需要先将传入的数组按 start 排序
*/
func merge(intervals [][]int) [][]int {
	length := len(intervals)
	var result interval
	sort.Sort(interval(intervals))
	if length < 2 {
		return intervals
	}
	temp := []int{intervals[0][0], intervals[0][1]}
	result = append(result, temp)
	i, j := 0, 1
	for j < length {
		if result[i][1] >= intervals[j][0] {
			result[i][0] = min(result[i][0], intervals[j][0])
			result[i][1] = max(result[i][1], intervals[j][1])
		} else {
			temp := []int{intervals[j][0], intervals[j][1]}
			result = append(result, temp)
			i++
		}
		j++
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 排序
type interval [][]int

func (in interval) Less(i, j int) bool {
	return in[i][0] < in[j][0]
}

func (in interval) Swap(i, j int) {
	in[i], in[j] = in[j], in[i]
}

func (in interval) Len() int {
	return len(in)
}

func main() {

}
