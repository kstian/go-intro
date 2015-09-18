// Write a program that can swap two integers (x := 1; y := 2; swap(&x, &y) should give you x=2 and y=1).
package main

import "fmt"

func main() {
	x := 2
	y := 5
	swap(&x,&y)
	fmt.Println("x : ", x)
	fmt.Println("y : ", y)
}

func swap(first *int, second *int) {
	temp := *first
	*first = *second
	*second = temp
}