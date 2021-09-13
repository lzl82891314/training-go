package main

import "fmt"

func main() {
	var num int
	num = 50

	ans1 := fibonacci_normal(num)
	fmt.Printf("the normal fibonacci func's answer with the n=[%v] is: [%v], and the TYPE of the answer is: [%T]\n", num, ans1, ans1)
	ans2 := fibonacci_memoized(num, make(map[int]int, num))
	fmt.Printf("the memoized fibonacci func's answer with the n=[%v] is: [%v], and the TYPE of the answer is: [%T]\n", num, ans2, ans2)
	ans3 := fibonacci_dp(num)
	fmt.Printf("the dp fibonacci func's answer with the n=[%v] is: [%v], and the TYPE of the answer is: [%T]\n", num, ans3, ans3)
}

func fibonacci_normal(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	return fibonacci_normal(n-1) + fibonacci_normal(n-2)
}

func fibonacci_memoized(n int, hash map[int]int) int {
	val, ok := hash[n]
	if ok {
		return val
	}
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	ans := fibonacci_memoized(n-1, hash) + fibonacci_memoized(n-2, hash)
	hash[n] = ans
	return ans
}

func fibonacci_dp(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	n1, n2 := 2, 3
	for i := 3; i < n; i++ {
		n2 += n1
		n1 = n2 - n1
	}
	return n1
}
