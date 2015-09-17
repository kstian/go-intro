//Using makeEvenGenerator as an example, write a makeOddGenerator function that generates odd numbers.

package main

import "fmt"

func main() {
	nextEven := makeOddGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4
}

func makeOddGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		if i == 0{
			i += 1
		}
		ret = i
		i += 2
		return
	}
}