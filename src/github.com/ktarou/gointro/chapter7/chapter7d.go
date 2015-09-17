/*
The Fibonacci sequence is defined as: fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2). Write a recursive function
which can find fib(n).
 */

package main

import "fmt"

func main() {
	fmt.Println(fib(9))
}
func fib(n uint) uint {
	if n ==0 {
		return 0
	}
	if n ==1 {
		return 1
	}
	return fib(n - 1) + fib(n - 2)
}