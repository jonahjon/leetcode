package main

// NumJewelsInStones function to tell how many jewels in stones
func NumJewelsInStones(J string, S string) int {
	isJewel := make(map[byte]bool, len(J))
	for i := range J {
		isJewel[J[i]] = true
	}

	res := 0
	for i := range S {
		if isJewel[S[i]] {
			res++
		}
	}

	return res
}
