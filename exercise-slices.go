package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    ry := make([][]uint8, dy)
    for i := range ry {
        ry[i] = make([]uint8, dx)
        for j := range ry[i] {
            ry[i][j] = uint8((i+j)/2)
        }
    }
    return ry
}

func main() {
    pic.Show(Pic)
}
