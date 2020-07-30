package main

import (
	"strconv"
)

//Tencent 2020 校招
//HG[3|B[2|CA]]F
func UnCompress(str string) string {

	var left, spilt, right int
	var cache string

	for i := range str {
		switch string(str[i]) {
		case "[":
			left = i
		case "|":
			spilt = i
		case "]":
			cache = ""
			right = i
			ar := str[spilt+1 : right]
			nm, _ := strconv.Atoi(str[left+1 : spilt])
			for nm != 0 {
				cache += ar
				nm--
			}
			str = str[:left] + cache + str[right+1:]
			str = UnCompress(str)
		}
	}
	return str
}
