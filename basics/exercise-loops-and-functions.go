package main

import "fmt"

func Sqrt(x float64) float64 {
	if x == 0 {
		return 0
	}
	z := 1.0
	for i := 0; i < int(x); i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	for i := 0; i < 15; i++ {
		fmt.Println(Sqrt(float64(i)))
	}
}
