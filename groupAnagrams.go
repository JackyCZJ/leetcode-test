package main

import "sort"

//Your runtime beats 67.75 % of golang submissions.
//Your memory usage beats 44.99 % of golang submissions.
func groupAnagrams(strs []string) [][]string {
	strsMap := make(map[string][]string)
	res := new([][]string)

	var sort = func(str string) string {
		var ii []int
		for i := range str {
			ii = append(ii, int(str[i]))
		}
		sort.Ints(ii)
		var i8 []uint8
		for i := range ii {
			i8 = append(i8, uint8(ii[i]))
		}
		return string(i8)
	}

	for i := range strs {
		index := sort(strs[i])
		strsMap[index] = append(strsMap[index], strs[i])
	}

	for i := range strsMap {
		*res = append(*res, strsMap[i])
	}

	return *res
}

