package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	pas := []int{1}
	for i := 0; i < n; i++ {
		fmt.Println(pas)
		pas = pascal(pas)
	}
}

func pascal(pas []int) []int {
	newpas := make([]int, len(pas)+1)
	newpas[0] = 1
	for i := 1; i < (len(newpas) - 1); i++ {
		newpas[i] = pas[i-1] + pas[i]
	}
	newpas[len(newpas)-1] = 1
	return newpas
}
