package compute

import "fmt"

func Add(nums []float64) (float64, error) {
	if (len(nums)) < 2 {
		return 0, fmt.Errorf("add requires at least 2 numbers")
	}

	sum := 0.0
	for _, x := range nums {
		sum += x
	}
	return sum, nil
}