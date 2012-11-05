package main

import (
    "flag"
    "fmt"
    "strconv"
    "math"
)

func isPrime(x float64) bool {
    if x <= 1 {
        return false
    }

    if math.Mod(x, 2) == 0 {
        return x == 2
    }

    for i := 3.0; i <= math.Sqrt(i); i += 2 {
        if math.Mod(x, i) == 0 {
            return false
        }
    }

    return true
}

func main() {
    flag.Parse()
    args := flag.Args()
    if len(args) < 1 || len(args) > 1 {
        fmt.Println("wronged")
        return
    }

    x, _ := strconv.ParseFloat(args[0], 64)
    fmt.Println(isPrime(x))
    return 1
}
