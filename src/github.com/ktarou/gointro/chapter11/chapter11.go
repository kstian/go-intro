package main
import (
	"github.com/ktarou/gointro/chapter11/math"
	"fmt"
)

func main()  {
	xs := []float64{1,2,3,4}
	avg := math.Average(xs)
	min := math.Min(xs)
	max := math.Max(xs)
	fmt.Println("Average : ", avg)
	fmt.Println("Maximum : ", max)
	fmt.Println("Minimum : ", min)
}
