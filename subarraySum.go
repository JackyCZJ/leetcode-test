package main

//Your runtime beats 74.47 % of golang submissions.
//Your memory usage beats 35.36 % of golang submissions.
func subarraySum(nums []int, k int) int {
	count:=0
	sum  := 0
	var m = make(map[int]int)
	for i := range nums{
		sum += nums[i]
		if sum == k {
			count++
		}
		count += m[sum-k]
		m[sum]++
	}
	return count
}

//Your runtime beats 14.96 % of golang submissions.
//Your memory usage beats 83.61 % of golang submissions.
//func subarraySum(nums []int, k int) int {
//	count:=0
//	for i := range nums{
//		sum := 0
//		for j := range nums[i:]{
//			sum += nums[i:][j]
//			if sum == k{
//				count++
//			}
//		}
//	}
//	return count
//}

