package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	num    []int
	target int
	ans    []int
}{

	{
		[]int{2, 7, 11, 15},
		9,
		[]int{0, 1},
	},
	{
		[]int{11, 15, 2, 7},
		9,
		[]int{2, 3},
	},
}

func Test_TwoSum(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, TwoSum(tc.num, tc.target), ":%v", tc)
	}
}
