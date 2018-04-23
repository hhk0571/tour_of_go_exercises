/* https://tour.golang.org/flowcontrol/8

Task: implement a square root function using Newton's method

*/

package main

import (
    "fmt"
)

const EPSILON = 0.000000000000001

func Sqrt(x float64) float64 {
    z := 1.0
    old_z := 0.0
    fmt.Println("========BEGIN==========")
    for i := 1; (old_z-z)*(old_z-z) > EPSILON; i++ {
        old_z = z
        z -= (z*z - x) / (2 * z)
        fmt.Println(i, ":", z)
    }
    fmt.Println("=========END===========")
    return z
}

func main() {
    fmt.Printf("Sqrt(%g) = %g\n", 2.0, Sqrt(2))
    fmt.Printf("Sqrt(%g) = %g\n", 3.0, Sqrt(3))
    fmt.Printf("Sqrt(%g) = %g\n", 5.0, Sqrt(5))

}

