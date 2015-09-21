package math
import "testing"

type test_pair struct {
	values []float64
	expected float64
}

var test_average = []test_pair{
	{ []float64{1,2}, 1.5 },
	{ []float64{1,1,1,1,1,1}, 1 },
	{ []float64{-1,1}, 0 },
}

var test_min = []test_pair{
	{ []float64{1,2}, 1 },
	{ []float64{4,1,5,7,1,4}, 1 },
	{ []float64{-1,1}, -1 },
	{ []float64{}, 0 },
}

var test_max = []test_pair{
	{ []float64{1,2}, 2 },
	{ []float64{4,1,5,7,1,4}, 7 },
	{ []float64{-1,1}, 1 },
	{ []float64{}, 0 },
}

func TestAverage(t *testing.T) {
	for _, pair := range test_average {
		v := Average(pair.values)
		if v != pair.expected {
			t.Error(
				"For", pair.values,
				"expected", pair.expected,
				"got", v,
			)
		}
	}
}

func TestMax(t *testing.T) {
	for _, pair := range test_max {
		v := Max(pair.values)
		if v != pair.expected {
			t.Error(
				"For", pair.values,
				"expected", pair.expected,
				"got", v,
			)
		}
	}
}

func TestMin(t *testing.T) {
	for _, pair := range test_min {
		v := Min(pair.values)
		if v != pair.expected {
			t.Error(
				"For", pair.values,
				"expected", pair.expected,
				"got", v,
			)
		}
	}
}