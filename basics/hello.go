package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println("Hello, world!")

// Values
	fmt.Println("go" + "lang")

    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)

// Variables
    var a = "initial"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    var e int
    fmt.Println(e)

    f := "apple"
    fmt.Println(f)

// Constants

	fmt.Println(s)

	const n = 5000000

	const g = 3e20 / n

	fmt.Println(g)

	fmt.Println(int64(g))

	fmt.Println(math.Sin(n))

// loops

 	i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }
	for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }
	for {
        fmt.Println("loop")
        break
    }
	for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }

// if/else

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}


}