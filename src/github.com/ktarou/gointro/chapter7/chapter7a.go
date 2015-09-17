/*
Write a function which takes an integer and halves it and returns true if it was even or false if it was odd.
For example half(1) should return (0, false) and half(2) should return (1, true).
*/

package main

import "fmt"

func main() {
	for num := 1; num <= 5 ; num++ {
		fmt.Println(half(num))
	}
}

func half(num int)(result int, even bool){
	result = num/2
	if num % 2 == 0 {
		even = true
	}
	return result, even
}
