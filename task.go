package main

import (
	"fmt"
	"math"
)

//#1
// func main() {
// 	var number1 int
// 	var base int
// 	fmt.Println("Введите число в десятичной системе")
// 	fmt.Scan(&number1)
// 	fmt.Println("Введите систему счисления")
// 	fmt.Scan(&base)
// 	if base < 2 || base > 36 {
// 		fmt.Println("не")
// 		return
// 	}
// 	fmt.Println("Ваше число равно:", perevod(number1, base))
// }

// func perevod(num int, base int) string {
// 	if num == 0 {
// 		return "0"
// 	}

// 	var result string
// 	for num > 0 {
// 		remainder := num % base
// 		var digit string

// 		if remainder < 10 {
// 			digit = string('0' + remainder)
// 		} else {
// 			digit = string('A' + remainder - 10)
// 		}

// 		result = digit + result
// 		num /= base
// 	}
// 	return result
// }

//#2
// func main() {
// 	var a, b, c float64
// 	fmt.Println("Введите коэффицент а")
// 	fmt.Scan(&a)
// 	fmt.Println("Введите коэффицент b")
// 	fmt.Scan(&b)
// 	fmt.Println("Введите коэффицент c")
// 	fmt.Scan(&c)
// 	root1, root2, found := find_roots(a, b, c)
// 	fmt.Println("Корни уравнения:", root1, ",", root2, "Корни есть?", found)
// }

// func find_roots(a float64, b float64, c float64) (float64, float64, bool) {
// 	d := b*b - 4*a*c
// 	if d == 0 {
// 		root := (-b + math.Sqrt(d)) / 2
// 		return root, root, true
// 	} else if d > 0 {
// 		root1 := (-b + math.Sqrt(d)) / 2
// 		root2 := (-b - math.Sqrt(d)) / 2
// 		return root1, root2, true
// 	} else {
// 		fmt.Println("Корней нет")
// 	}
// 	return 0, 0, false
// }

// #3
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

// #4
// func main() {
// 	arr_main := []int{1, 4, 2, 4}
// 	fmt.Println(arr_main)
// 	arr1 := []int{1, 6, 3, 4, 5}
// 	arr_main = arr_append(arr_main, arr1)
// 	arr_main = arr_sort(arr_main)
// 	fmt.Println(arr_main)
// }

// func arr_append(arr_main []int, arr_to_append []int) []int {
// 	arr_sort(arr_main)
// 	arr_sort(arr_to_append)
// 	fmt.Println(arr_main, arr_to_append)
// 	for i := 0; i < len(arr_to_append); i++ {
// 		arr_main = append(arr_main, arr_to_append[i])
// 	}
// 	return arr_main
// }

// func arr_sort(arr []int) []int {
// 	for i := 0; i < (len(arr) - 1); i++ {
// 		for j := 0; j < ((len(arr) - i) - 1); j++ {
// 			if arr[j] > arr[j+1] {
// 				arr[j], arr[j+1] = arr[j+1], arr[j]
// 			}
// 		}
// 	}
// 	return arr
// }

// #5
// func main() {
// 	a := "vsem privet"
// 	b := "priv"
// 	index := find_str(a, b)
// 	if index == -1 {
// 		fmt.Println("Не найдено")
// 	} else {
// 		fmt.Printf("Подстрока '%s', найдена в строке '%s', начиная с '%d' индекса", b, a, index)
// 	}
// }

// func find_str(a, b string) int {
// 	var flag bool
// 	alen := len(a)
// 	blen := len(b)
// 	if alen < blen {
// 		return -1
// 	}
// 	for i := 0; i < alen; i++ {
// 		flag = true
// 		for j := 0; j < blen; j++ {
// 			if a[i+j] != b[j] {
// 				flag = false
// 				break
// 			}
// 		}
// 		if flag {
// 			return i
// 		}
// 	}
// 	return -1
// }

// #5
// func main() {
// 	a := "huhqqq"
// 	fmt.Println(palindrom(a))
// }

// func palindrom(a string) bool {
// 	left := 0
// 	right := len(a) - 1
// 	for left < right {
// 		if a[left] != a[right] {
// 			return false
// 		}
// 		left++
// 		right--
// 	}
// 	return true
// }

// #calc
// func main() {
// 	var a, b float64
// 	var operation string
// 	fmt.Println("Введите первое число: ")
// 	fmt.Scan(&a)
// 	fmt.Println("Введите операцию: ")
// 	fmt.Scan(&operation)
// 	fmt.Println("Введите второе число: ")
// 	fmt.Scan(&b)
// 	switch operation {
// 	case "+":
// 		fmt.Println("Результат сложения: ", a+b)
// 	case "-":
// 		fmt.Println("Результат вычитания: ", a-b)
// 	case "*":
// 		fmt.Println("Результат умножения: ", a*b)
// 	case "/":
// 		if b != 0 {
// 			fmt.Println("Результат деления: ", a/b)
// 		} else {
// 			fmt.Println("На ноль делить нельзя!")
// 		}
// 	case "^":
// 		fmt.Println("Результат возведения в степерь: ", math.Pow(a, b))
// 	case "%":
// 		var x int = int(a)
// 		var y int = int(b)
// 		fmt.Println("Результат поиска отстатка от деления:", x%y)
// 	default:
// 		fmt.Println("Выбранная операция неккоректна")
// 	}
// }

// #3
// func main() {
// 	x := [2]int{101, 150}
// 	y := [2]int{99, 100}
// 	z := [2]int{14, 26}
// 	fmt.Println(find_for_three(x, y, z))

// }

// func find_for_two(x, y [2]int) bool {
// 	return x[0] <= y[1] && y[0] <= x[1]
// }

// func find_for_three(x, y, z [2]int) bool {
// 	return find_for_two(x, y) || find_for_two(x, z) || find_for_two(y, z)
// }
