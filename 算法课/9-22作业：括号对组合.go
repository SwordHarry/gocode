package main

import "fmt"

/**
给出n，打印出n对合法的括号对组合，例如n=3，输出如下：
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"

*/

func printBrackets(n int) {
	printStrWithBrackets(n, n, "")
}

// 递归：左括号 和 右括号
func printStrWithBrackets(leftB, rightB int, s string) {
	if leftB == rightB && leftB == 0 {
		fmt.Print(s, " ")
	}
	if leftB > 0 {
		printStrWithBrackets(leftB-1, rightB, s+"(")
	}
	if leftB < rightB {
		printStrWithBrackets(leftB, rightB-1, s+")")
	}
}

func main() {
	fmt.Println("n = 2:")
	printBrackets(2)
	fmt.Println()
	fmt.Println("n = 3:")
	printBrackets(3)
	fmt.Println()
	fmt.Println("n = 4:")
	printBrackets(4)
	fmt.Println()
}
