package main

/*
Your runtime beats 48.86 % of golang submissions.
Your memory usage beats 31.21 % of golang submissions.
func singleNumber(nums []int) int {
	var numMap = make(map[int]bool, len(nums))
	for n := range nums {
		if _, ok := numMap[nums[n]]; ok {
			numMap[nums[n]] = false
			continue
		} else {
			numMap[nums[n]] = true
		}
	}
	for n := range numMap {
		if numMap[n] == true {
			return n
		}
	}
	return 0
}
*/
//xor way
//Copy , not my solution
//Runtime: 8 ms, faster than 95.55% of Go online submissions for Single Number.
//Memory Usage: 4.7 MB, less than 100.00% of Go online submissions for Single Number.
func singleNumber(nums []int) int {
	for i := 1 ; i < len(nums);i ++{
		nums[0] ^= nums[i]
	}
	return nums[0]
}