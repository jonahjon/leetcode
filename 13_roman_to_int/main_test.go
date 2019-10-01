package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	str string
	ans int
}{

	{
		"III",
		3,
	},

	{
		"IV",
		4,
	},
	{
		"IX",
		9,
	},
	{
		"LVIII",
		58,
	},
}

func Test_RomanToInt(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, RomanToInt(tc.str), ":%d", tc)
	}
}
