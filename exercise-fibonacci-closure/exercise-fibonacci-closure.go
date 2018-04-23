/* https://tour.golang.org/moretypes/26

Task: Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...)
*/

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    a, b := 0, 1
    return func() int {
        ret := a
        a, b = b, a+b
        return ret
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 20; i++ {
        fmt.Print(f(), " ")
    }
}


