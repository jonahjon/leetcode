package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	str string
	ans string
}{

	{
		"(()())(())",
		"()()()",
	},
	{
		"(()())(())(()(()))",
		"()()()()(())",
	},
	{
		"()()",
		"",
	},
}

func Test_numJewelsInStones(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, RemoveOuterParentheses(tc.str), ":%v", tc)
	}
}
