package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {

    z := 2.0
    oldz := 0.0

    for math.Abs(z - oldz) > 0.000001 {
        oldz = z
        z = z - ((z * z) - x) / (2 * z)
    }
    return z
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(4))
    fmt.Println(Sqrt(9))
}
