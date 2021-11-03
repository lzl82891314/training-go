package slice

import (
	"math/rand"
	"time"
)

//func main() {
//	num := make([]int, 100, 100)
//	InitArray(num)
//	fmt.Println(num)
//	BubbleSort(num)
//	fmt.Println(num)
//}

func BubbleSort(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if arr[i] < arr[j] {
				tmp := arr[i]
				arr[i] = arr[j]
				arr[j] = tmp
			}
		}
	}
}

func InitArray(array []int) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(100)
	}
}
