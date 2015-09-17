package main
import "fmt"

const ft = 0.3048
var input float64
func main() {
	fmt.Print("Enter value in feet(ft) : ")
	fmt.Scanf("%f", &input)//put the value into input memory address, note the & sign
	result := input * ft
	fmt.Println("Result in meter is : ", result)
}
