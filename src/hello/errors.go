package main

import (
	"math"
	"fmt"
)

type NegativeSqrtError float64

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, NegativeSqrtError(f)
	}

	return math.Sqrt(f), nil
}


func (f NegativeSqrtError) Error() string {
	return fmt.Sprintf("math: square root of negative number %g", float64(f))
}

func main() {
	f, err := Sqrt(-1)

	v, ok := err.(NegativeSqrtError)
	if ok {
		fmt.Println(err)
		fmt.Println(v)
	} else {
		fmt.Println(f)
	}
}
