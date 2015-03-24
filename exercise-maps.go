package main

import (
    "golang.org/x/tour/wc";
    "strings";
)

func WordCount(s string) map[string]int {
    a := []string
    a = strings.Fields(s)
    m := make(map[string]int)
    for _,value := range a {
        _, ok := m[value]
        if ok {
            m[value]++
        } else {
            m[value] = 1
        }
    }
    return m
}

func main() {
    wc.Test(WordCount)
}

