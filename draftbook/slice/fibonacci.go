package slice

//func main() {
//	var num int
//	num = 25
//
//	ans1 := Fibonacci_normal(num)
//	fmt.Printf("the normal fibonacci func's answer with the n=[%v] is: [%v], and the TYPE of the answer is: [%T]\n", num, ans1, ans1)
//	hash := make(map[int]int, num)
//	ans2 := Fibonacci_memoized01(num, hash)
//	fmt.Printf("the memoized fibonacci func's answer with the n=[%v] is: [%v], and the TYPE of the answer is: [%T]\n", num, ans2, ans2)
//	ans3 := Fibonacci_dp01(num)
//	fmt.Printf("the dp fibonacci func's answer with the n=[%v] is: [%v], and the TYPE of the answer is: [%T]\n", num, ans3, ans3)
//	ans4 := Fibonacci_dp02(num)
//	fmt.Printf("the dp2 fibonacci func's answer with the n=[%v] is: [%v], and the TYPE of the answer is: [%T]\n", num, ans4, ans4)
//}

func FibonacciNormal(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	return FibonacciNormal(n-1) + FibonacciNormal(n-2)
}

func FibonacciMemoized01(n int, hash map[int]int, arr [10]int) int {
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
	for i := 0; i < len(arr); i++ {
		arr[i] = i * n
	}
	hash[n] = FibonacciMemoized01(n-1, hash, arr) + FibonacciMemoized01(n-2, hash, arr)
	return hash[n]
}

func FibonacciMemoized02(n int, hash *map[int]int, arr *[10]int) int {
	val, ok := (*hash)[n]
	if ok {
		return val
	}
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	for i := 0; i < len(arr); i++ {
		arr[i] = i * n
	}
	(*hash)[n] = FibonacciMemoized02(n-1, hash, arr) + FibonacciMemoized02(n-2, hash, arr)
	return (*hash)[n]
}

func FibonacciDp01(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	hash := make(map[int]int, n)
	hash[1] = 1
	hash[2] = 1
	for i := 3; i <= n; i++ {
		hash[i] = hash[i-1] + hash[i-2]
	}
	return hash[n]
}

func FibonacciDp02(n int) int {
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
