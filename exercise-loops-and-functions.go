package main

import (
    "fmt";
    "math";
)

func Sqrt(x float64, y int) float64 {
    z := float64(1)
    for i:=0;i<y;i++ {
        z = z - (z*z - x)/(2*z)
    }
    return z
}

func main() {
    i := 1
    z_prev:=Sqrt(2,i)
    i+=1
    for z_cur:=Sqrt(2,i); z_cur!=z_prev; i++ {
        z_prev=z_cur
    }
    fmt.Println(z_prev)
    fmt.Println(math.Sqrt(2))
    fmt.Println(math.Sqrt(2))
}
