// Write a function with one variadic parameter that finds the greatest number in a list of numbers.

package main

import "fmt"

func main() {
	fmt.Println(findGreatest(5, 4, 7, 3, 9, 10))
}

func findGreatest(num... int) (max int) {
	for i, value := range num {
		if i == 0 {
			max = value
		}
		if value > max {
			max = value
		}
	}
	return
}