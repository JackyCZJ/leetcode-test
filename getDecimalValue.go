package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	var arr []int
	for {
		arr = append(arr, head.Val)
		if head.Next != nil {
			head = head.Next
		} else {
			break
		}
	}
	i, j, s := len(arr)-1, 0, 0
	for i >= 0 {
		s += (1 << j) * arr[i]
		i--
		j++
	}
	return s
}
