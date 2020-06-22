package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "first",
			args: args{strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			want: [][]string{{"ate","eat","tea"},{"nat","tan"},{"bat"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got :=  groupAnagrams(tt.args.strs)
			resMap := make(map[int]bool)
			for i := range got {
				resMap[i] = false
				for j :=range tt.want{
					if arraySortedEqual(got[i],tt.want[j]){
						resMap[i] = true
					}
				}
			}
			for i :=range resMap{
				if !resMap[i]{
					t.Error("error")
				}
			}
		})
	}
}

func arraySortedEqual(a, b []string) bool {
	if len(a) != len(b) {return false }

	aCopy := make([]string, len(a))
	bCopy := make([]string, len(b))

	copy(aCopy, a)
	copy(bCopy, b)

	sort.Strings(aCopy)
	sort.Strings(bCopy)

	return reflect.DeepEqual(aCopy, bCopy)
}