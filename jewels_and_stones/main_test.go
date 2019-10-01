package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test Cases
var tcs = []struct {
	J   string
	S   string
	ans int
}{

	{
		"aA",
		"aAAbbbb",
		3,
	},

	{
		"z",
		"ZZ",
		0,
	},
}

func Test_numJewelsInStones(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, NumJewelsInStones(tc.J, tc.S), "输入:%v", tc)
	}
}
