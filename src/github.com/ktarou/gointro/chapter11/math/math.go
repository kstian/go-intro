package math
// Finds the average of a series of numbers
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}
// Finds the maximum of a series of numbers
func Max(xs []float64) (max float64) {
	for i, value := range xs {
		if i == 0 {
			max = value
		}
		if value > max {
			max = value
		}
	}
	return
}
// Finds the minimum of a series of numbers
func Min(xs []float64) (min float64) {
	for i, value := range xs {
		if i == 0{
			min = value
		}
		if value < min{
			min = value
		}
	}
	return
}
