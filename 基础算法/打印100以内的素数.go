package main

import "fmt"

/**
素数指只能被 1 和其自身整除的正数

思路：数学遍历，外层 for 循环遍历100以内正数，内层 for 循环遍历除数
优化：所有偶数除2以外，都能被2整除，则 i+=2 即可；
	内层循环从 3 循环到 j*j 即可，因为对于一个数 n，若存在可除数，则要么为平方，要么一大数乘一小数
*/

func getAllNum(num int) []int {
	var result []int
	for i := 3; i <= num; i += 2 {
		flag := true
		for j := 3; j*j <= i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, i)
		}
	}
	return result
}

func main() {

	fmt.Println(getAllNum(100))
}
