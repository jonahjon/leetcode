package main

import "strings"

// NumJewelsInStones function to tell how many jewels in stones
func DefangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", -1)
}
