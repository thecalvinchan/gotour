package main

import (
    "fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64, y int) (float64, error) {
    if (x < 0) {
        return x, ErrNegativeSqrt(x)
    }
    z := float64(1)
    for i:=0;i<y;i++ {
        z = z - (z*z - x)/(2*z)
    }
    return z, nil
}

func main() {
    fmt.Println(Sqrt(2,10))
    fmt.Println(Sqrt(-2,10))
}

