package slice

import (
	"testing"
)

var num int = 35

func BenchmarkFibonacciNormal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciNormal(num)
	}
}

func BenchmarkFibonacciMemoized01(b *testing.B) {
	arr := [10]int{1, 3, 4}
	for i := 0; i < b.N; i++ {
		hash := make(map[int]int, num)
		FibonacciMemoized01(num, hash, arr)
	}
	// fmt.Println(arr)
}

func BenchmarkFibonacciMemoized02(b *testing.B) {
	arr := [10]int{1, 3, 4}
	for i := 0; i < b.N; i++ {
		hash := make(map[int]int, num)
		FibonacciMemoized02(num, &hash, &arr)
	}
	// fmt.Println(arr)
}

func BenchmarkFibonacciDp01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciDp01(num)
	}
}

func BenchmarkFibonacciDp02(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciDp02(num)
	}
}
