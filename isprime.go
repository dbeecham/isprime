package main

import (
    "flag"
    "fmt"
    "math"
    "strconv"
    "os"
)

// This function will eventually get it's own module.
func isPrime(t int64) bool {

    // math.Mod requires floats.
    x := float64(t)

    // 1 or less aren't primes.
    if x <= 1 {
        return false
    }

    // Solve half of the integer set directly
    if math.Mod(x, 2) == 0 {
        return x == 2
    }

    // Main loop. i needs to be float because of math.Mod.
    for i := 3.0; i <= math.Floor(math.Sqrt(x)); i += 2.0 {
        if math.Mod(x, i) == 0 {
            return false
        }
    }

    // It's a prime!
    return true
}

func main() {
    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "usage: %s int\n", os.Args[0])
        flag.PrintDefaults()
        os.Exit(2)
    }

    flag.Parse()

    args := flag.Args()
    if len(args) < 1 || len(args) > 1 {
        flag.Usage()
    }

    // I really shouldn't omit the error message here...
    x, _ := strconv.ParseInt(args[0], 0, 64)

    // Should i os.exit(isPrime())?
    // It's more logical to do so, but not convenient for most shell users.
    fmt.Println(isPrime(x))
}
