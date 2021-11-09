// https://go-tour-jp.appspot.com/flowcontrol/8
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z float64 = 1
	count := 0
	for {
		result := z - (z*z - x) / (2 * z)
		if math.Abs(z-result) < 0.001 {
			break
		}
		count += 1
		z = result
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}