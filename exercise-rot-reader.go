package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}


func (reader rot13Reader) Read(b []byte) (int, error) {
    i, err := reader.r.Read(b);
    for j:=0;j<i;j++ {
        if (b[j] >= 'A' && b[j] < 'N') || (b[j] >='a' && b[j] < 'n') {
          b[j] += 13
    
        // if the letter's index is between M - Z, subtract 13 from its index
        } else if (b[j] > 'M' && b[j] <= 'Z') || (b[j] > 'm' && b[j] <= 'z'){
          b[j] -= 13
        }
    }
    return i, err
}

func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
