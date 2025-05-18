// Задача 1
//
// //package main
//
//import (
//   "fmt"
//   "math"
//)
//
//func main() {
// Задаем значения a и b и x
//    fmt.Println(calculate(7.2, 4.2,  2.4))
//}
//func calculate(a, b, x float64) float64 {
//   numerator := math.Abs(a - b*x)
//
//   denominator := math.Pow(math.Log10(x), 3)
//
//   result := math.Sqrt(numerator / denominator)
//
// return result
//}

//Задача 2

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(calculate(7.2, 4.2, 2.4))

	var arr [5]float64 = [5]float64{2.4, 2.8, 3.9, 4.7, 3.16}
	fmt.Println(arr)
	for _, el := range arr {
		fmt.Println(calculate(7.2, 4.2, el))
	}
}
func calculate(a, b, x float64) float64 {
	numerator := math.Abs(a - b*x)

	denominator := math.Pow(math.Log10(x), 3)

	result := math.Sqrt(numerator / denominator)

	return result
}