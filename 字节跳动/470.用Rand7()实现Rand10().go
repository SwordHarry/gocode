package main

import "math/rand"

/**
已有方法rand7可生成 1 到 7 范围内的均匀随机整数，试写一个方法rand10生成 1 到 10 范围内的均匀随机整数。

不要使用系统的Math.random()方法。

示例 1:
输入: 1
输出: [7]

示例 2:
输入: 2
输出: [8,4]

示例 3:
输入: 3
输出: [8,1,10]

提示:
rand7已定义。
传入参数:n表示rand10的调用次数。

进阶:
rand7()调用次数的期望值是多少?
你能否尽量少调用 rand7() ?
*/

/**
思路1：拒绝采样
	1 - 42 均匀生成，舍弃 41 和 42
	用 40
*/
func rand10() int {
	x, y := rand7(), rand7()
	z := (x-1)*7 + y // 均匀生成 1- 42
	for ; z > 40; z = (x-1)*7 + y {
		x, y = rand7(), rand7()
	}
	return 1 + (z-1)%10 // (z-1)%10 为了均匀生成 0-9
}

/**
思路2：思路1 基础上缩减拒绝采样的范围
*/
func rand10Plus() int {
	x, y := rand7(), rand7()
	z := (x-1)*7 + y // 均匀生成 1- 42
	for {
		x, y = rand7(), rand7()
		z = (x-1)*7 + y
		if z <= 40 {
			return 1 + (z-1)%10
		}
		z = (z-40-1)*rand7() + rand7()
		if z <= 60 {
			return 1 + (z-1)%10
		}
		z = (z-60-1)*rand7() + rand7()
		if z <= 20 {
			return 1 + (z-1)%10
		}
	}
}

func rand7() int {
	return rand.Intn(7)
}

func main() {

}
