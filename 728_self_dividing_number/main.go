package main

import "fmt"

// SelfDividingNumbers the function to lower case values
func SelfDividingNumbers(left, right int) []int {
	res := make([]int, 0, right-left+1)
	for i := left; i <= right; i++ {
		if selfdividing(i) {
			res = append(res, i)
		}
	}
	fmt.Println(res)
	return res
}

func selfdividing(n int) bool {
	t := n
	for t > 0 {
		d := t % 10
		t /= 10
		if d == 0 || n%d != 0 {
			return false
		}
	}
	return true
}
