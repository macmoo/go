package main

import (
	"fmt"
	"math"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("제곱근은 양수여야 한다. f:%g", f)
	}
	return math.Sqrt(f), nil
}

func main() {
	fmt.Println("---------------------------------")
	sqrt, err := Sqrt(-2)
	if err != nil {
		fmt.Printf("Error:%v\n", err)
		return
	}
	fmt.Printf("Sqrt(-2)=%v\n", sqrt)
}

// ---------------------------------
// Error:제곱근은 양수여야 한다. f:-2
