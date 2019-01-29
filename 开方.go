package main

import (
	"fmt"
	"math"
)

//牛顿法实现开方运算

func Sqrt(x float64) float64 {
	var z float64 = 1.0
	for i := 1; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
		fmt.Println(z)
	}

	return z
}

func main() {
	fmt.Printf("My sqrt(%d) is %g\n", 3, Sqrt(3))
	fmt.Printf("math.Sqrt(%d) is %g\n", 3, math.Sqrt(3))
}
