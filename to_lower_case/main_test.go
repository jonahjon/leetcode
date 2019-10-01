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
		"Hello",
		"hello",
	},

	{
		"Yolo",
		"yolo",
	},
}

func Test_numJewelsInStones(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, ToLowerCase(tc.str), "输入:%v", tc)
	}
}
