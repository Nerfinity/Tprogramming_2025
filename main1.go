package main

import (
    "fmt"
    "math"
)

func main() {
    // Задаем значения a и b и x
    fmt.Println(calculate(7.2, 4.2,  2.4))
}
func calculate(a, b, x float64) float64 {
    numerator := math.Abs(a - b*x)

    denominator := math.Pow(math.Log10(x), 3)

    result := math.Sqrt(numerator / denominator)

    return result
}