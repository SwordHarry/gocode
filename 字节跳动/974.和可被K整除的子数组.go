package main

/**
给定一个整数数组 A，返回其中元素之和可被 K 整除的（连续、非空）子数组的数目。
示例：
输入：A = [4,5,0,-2,-3,1], K = 5
输出：7
解释：
有 7 个子数组满足其元素之和可被 K = 5 整除：
[4, 5, 0, -2, -3, 1], [5], [5, 0], [5, 0, -2, -3], [0], [0, -2, -3], [-2, -3]

提示：
1 <= A.length <= 30000
-10000 <= A[i] <= 10000
2 <= K <= 10000
*/

/**
思路：原数组 -> 前缀和数组 -> 求模数组 -> 哈希 -> 组合c(m,n)
	涉及到数学公式： (x-y)%n == 0 <-> x % n == y % n
	并且负数取模要保证正数：x = (x%n + n)%n
*/
func subarraysDivByK(A []int, K int) int {
	length := len(A)
	var result int
	if length == 0 {
		return result
	}

	// 前缀和数组
	// (x-y)%n == 0 -> x%n==y%n
	for i := 1; i < length; i++ {
		A[i] += A[i-1]
	}
	m := map[int]int{0: 1}
	for i := 0; i < length; i++ {
		// 出现负数需要纠正，因为负数的取模运算不统一
		// 注意 [2,-2,2,-4] K=6 的个例
		A[i] = (A[i]%K + K) % K
		m[A[i]]++
	}
	// 组合问题
	for _, v := range m {
		if v > 1 {
			result += v * (v - 1) / 2
		}
	}

	return result
}

func main() {

}
