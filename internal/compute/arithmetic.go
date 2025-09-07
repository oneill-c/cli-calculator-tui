package compute

import (
	"fmt"
)

func numsLengthGuard(nums []float64) error {
	if (len(nums)) < 2 {
		return fmt.Errorf("add requires at least 2 numbers")
	}
	return nil
}

func Add(nums []float64) (float64, error) {
	err := numsLengthGuard(nums)
	if err != nil {
		return 0, err
	}

	sum := 0.0
	for _, x := range nums {
		sum += x
	}
	return sum, nil
}

func Sub(nums []float64) (float64, error) {
	err := numsLengthGuard(nums)
	if err != nil {
		return 0, err
	}

	sum := nums[0]
	for _, x := range nums[1:] {
		sum -= x
	}
	return sum, nil
}

func Mul(nums []float64) (float64, error) {
	err := numsLengthGuard(nums)
	if err != nil {
		return 0, err
	}

	sum := 1.0
	for _, x := range nums {
		sum *= x
	}

	return sum, nil
}

func Div(a float64, b float64) (float64, error) {
	fmt.Println("I am here")

	fmt.Println(a)
	fmt.Println(b)


	if b == 0 {
		return 0, fmt.Errorf("cannot divide by 0")
	}
	return a / b, nil
}