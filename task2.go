package main

import (
	"fmt"
	"strconv"
)

// func main() {
// 	var s, v_0, v, a, t float64
// 	fmt.Println("Введите: начальную скорость, скорость, ускорение, время")
// 	fmt.Scan(&v_0)
// 	fmt.Scan(&v)
// 	fmt.Scan(&a)
// 	fmt.Scan(&t)
// 	if t < 0 {
// 		fmt.Println("Время не может быть отрицательным")
// 	} else {
// 		s = v_0*t + 0.5*((a*t*t)/2)
// 	}
// 	fmt.Println(s)
// }

//#3
// func main() {
// 	var x int
// 	fmt.Scan(&x)
// 	if visoko(x) == true {
// 		fmt.Println("Высокосный год")
// 	} else {
// 		fmt.Println("хуйня")
// 	}
// }

// func visoko(x int) bool {
// 	if (x%4 == 0 && x%100 != 0) || (x%400 == 0) {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// func main() {
// 	var x int
// 	fmt.Scan(&x)
// 	if Is_prime(x) == true {
// 		fmt.Println("Число простое")
// 	} else {
// 		fmt.Println("Число составное")
// 	}
// }

// func Is_prime(x int) bool {
// 	y := float64(x)
// 	x = int(math.Sqrt(y))
// 	for i := 2; i < x; i++ {
// 		if x%i == 0 {
// 			return false
// 		} else {
// 			return true
// 		}
// 	}
// 	return true
// }

func main() {
	var x int
	fmt.Scan(&x)
	var num_arr []int
	numstr := strconv.Itoa(x)
	for _, n := range numstr {
		num_arr = append(num_arr, strconv.Atoi(string(n)))
	}
}

// func amstrong_num(x int) int {
// 	numstr := strconv.Itoa(x)
// 	for i := range numstr {
// 		fmt.Println(numstr[i])
// 	}
// }

// func nums_arr(x int) []int {
// 	var numbers []int
// 	for x > 0 {
// 		numbers = append(numbers, (x % 10))
// 		x /= 10
// 	}
// 	numbers = numbers[]
// }
