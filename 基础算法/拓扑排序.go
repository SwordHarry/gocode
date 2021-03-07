package main

import "fmt"

/*
{
  "algorithms" : ["data structures"],
  "calculus": ["linear algebra"],
  "compliers": [
       "data structures",
       "formal languages",
       "computer organization",
     ],
 "data structures":["mathematics"],
 }
// 以上的map中的元组key是一门学科名，value是学科的列表，表示key这门学科依赖列表里的所有学科（先学value列表里的学科后，才能学key这门学科）
// 请打印出一种要学习所有学科的先后顺序列表；编程答题时间：半个小时到1个小时。
*/

/**
思路：拓扑排序 dfs，需要使用到 一个辅助哈希表 和 dfs 递归函数
*/

type searchType int

const (
	beforeSearch searchType = iota
	searching
	afterSearch
)

func getListOfSubject(m map[string][]string) []string {
	var (
		result []string
		sub    = make(map[string]searchType)
		dfs    func(course string)
		valid  = true
	)

	dfs = func(course string) {
		sub[course] = searching
		// 先将依赖全学完，如果没有依赖，直接该课程为 afterSearch
		for _, c := range m[course] {
			if sub[c] == beforeSearch {
				dfs(c)
				if !valid {
					return
				}
			} else if sub[c] == searching {
				valid = false
				return
			}
		}
		sub[course] = afterSearch
		result = append(result, course)
	}
	for k := range m {
		if valid && sub[k] == beforeSearch {
			dfs(k)
		} else if !valid {
			return nil
		}
	}
	return result
}

func main() {
	fmt.Println(getListOfSubject(map[string][]string{
		"a": {
			"b",
		},
		"c": {
			"d",
		},
		"e": {
			"b",
			"f",
			"g"},
		"b": {
			"h",
		},
	}))
}
