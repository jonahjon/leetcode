package main

//TwoSum takes the array of ints and returns an numbers that up to a specific target
func TwoSum(nums []int, target int) []int {
	index := make(map[int]int, len(nums))
	for i, b := range nums {
		if j, ok := index[target-b]; ok {
			return []int{j, i}
		}
		index[b] = i
	}
	return nil
}
