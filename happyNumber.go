package main

//Your memory usage beats 79.48 % of golang submissions.
//Your runtime beats 100.00 % of golang submissions.
func isHappy(n int) bool {
	digitSquareSum := func(num int) int{
		var sum, tmp int
		for num != 0{
			tmp = num % 10
			sum += tmp * tmp
			num /= 10
		}
		return sum
	}

	slow, fast :=  n , n
	for{
		slow = digitSquareSum(slow)
		fast = digitSquareSum(fast)
		fast = digitSquareSum(fast)
		if fast == 1{
			return true
		}
		if slow == fast{
			return false
		}
	}
}