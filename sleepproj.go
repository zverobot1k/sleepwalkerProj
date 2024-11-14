package main

import (
	"fmt"
)

// func main() {
// 	numbers := []int{10, -9, -40, 35, 90}
// 	fmt.Println("Массив до сортировки:", numbers)
// 	sort(numbers)
// 	fmt.Println("Массив после сортировки по абсолютным значениям:", numbers)
// }

// func sort(arr []int) {
// 	for i := 0; i < (len(arr) - 1); i++ {
// 		if math.Abs(float64(arr[i])) > math.Abs(float64(arr[i+1])) {
// 			arr[i], arr[i+1] = arr[i+1], arr[i]
// 		}
// 	}
// }

func main() {
	var a int
	var b int
	fmt.Scan(&a)
	fmt.Scan(&b)
}

func summ(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}
