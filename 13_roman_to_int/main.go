package main

//RomanToInt the function to lower case values
func RomanToInt(s string) int {
	romanMap := map[string]int{}
	romanMap["I"] = 1
	romanMap["V"] = 5
	romanMap["X"] = 10
	romanMap["L"] = 50
	romanMap["C"] = 100
	romanMap["D"] = 500
	romanMap["M"] = 1000

	MapIndex := string(s[len(s)-1])
	sum := romanMap[MapIndex]
	for i := len(s) - 2; i >= 0; i-- {
		if romanMap[string(s[i])] < romanMap[string(s[i+1])] {
			sum -= romanMap[string(s[i])]
		} else {
			sum += romanMap[string(s[i])]
		}
	}
	return sum
}
