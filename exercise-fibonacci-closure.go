package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    var fib int
    var fib2 int
    return func() int {
        if fib == 0 {
            fib = 1
        } else if fib2 == 0 {
            fib2 = fib
            fib = 1
        } else {
            fib = fib + fib2
            fib2 = fib - fib2
        }
        return fib
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}

