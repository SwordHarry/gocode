package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	flag.IntVar(&n, "n", 0, "n")
	flag.IntVar(&length, "m", 1, "m")
	flag.Parse()
	c(n, length)
}

var a = [100]int{}
var n, length int

// c(n,m) = c(n - 1, m) + c(n - 1,m - 1)
func c(n, m int) {
	for i := n; i >= m; i-- {
		a[m-1] = i
		if m == 1 {
			show(length)
		} else {
			c(i-1, m-1)
		}
	}
}

func show(length int) {
	for i := 0; i < length; i++ {
		fmt.Print(strconv.Itoa(a[i]) + " ")
	}
	fmt.Println()
}
