package main

//Your runtime beats 93.14 % of golang submissions.
//Your memory usage beats 67.46 % of golang submissions.
func twoSumI(nums []int, target int) []int {
	var numsMap = make(map[int]int, len(nums))
	for i := range nums {
		_, ok := numsMap[target-nums[i]]
		if ok {
			return []int{numsMap[target-nums[i]], i}
		}

		numsMap[nums[i]] = i
	}
	return []int{}
}

//Your runtime beats 95.79 % of golang submissions.
//Your memory usage beats 71.48 % of golang submissions.
func twoSumII(numbers []int, target int) []int {
	var i, j = 0, len(numbers) - 1
	var sum int
	for {
		sum = numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		}
		if sum > target {
			j--
		} else {
			i++
		}
	}
}


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//twoSum IV
//Your runtime beats 70.64 % of golang submissions
//Your memory usage beats 12.99 % of golang submissions.
func findTarget(root *TreeNode, k int) bool {
	if root.Right == nil && root.Left == nil {
		return false
	}

	var numbers []int
	var find func(i *TreeNode, a *[]int)
	find = func(i *TreeNode, a *[]int) {
		if i != nil {
			find(i.Left, a)
			*a = append(*a, i.Val)
			find(i.Right, a)
		}
	}
	find(root, &numbers)

	var i, j = 0, len(numbers) - 1
	var sum int
	for i < j {
		sum = numbers[i] + numbers[j]
		if sum == k {
			return true
		}
		if sum > k {
			j--
		} else {
			i++
		}
	}
	return false
}

