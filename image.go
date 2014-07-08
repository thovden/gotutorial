package main

import (
    "code.google.com/p/go-tour/pic"
    "fmt"
)

func Pic(dx, dy int) [][]uint8 {
    fmt.Println(dx, dy)
    
    yp := make([][]uint8, dy, dy)
    
    for i, _ := range yp {
        xp := make([]uint8, dx, dx)
        yp[i] = xp
        for j, _ := range xp {
        	yp[i][j] = uint8(dx*dy)
        }
    }
    return yp
}

func main() {
    pic.Show(Pic)
}
