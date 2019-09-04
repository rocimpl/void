package main

import (
    "fmt"
    "time"
)

func main() {
    for {
        t := time.Now()
        fmt.Printf("%v %d\n", t, int16(t.UnixNano()  % 1e9 / 1e5))
    }
}
