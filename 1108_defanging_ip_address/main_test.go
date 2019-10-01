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
		"1.1.1.1",
		"1[.]1[.]1[.]1",
	},

	{
		"255.100.50.0",
		"255[.]100[.]50[.]0",
	},
}

func Test_numJewelsInStones(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, DefangIPaddr(tc.str), ":%v", tc)
	}
}
