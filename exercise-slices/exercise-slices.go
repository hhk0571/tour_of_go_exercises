/* https://tour.golang.org/moretypes/18

Task: Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers. When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.

The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.

*/

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    z := make([][]uint8, dy)

    for i := range z {
        z[i] = make([]uint8, dx)
    }

    for i := range z {
        for j := range z[i] {
            z[i][j] = uint8(i*j)
            /* you can try other different computings.
            z[i][j] = uint8(i ^ j)
            z[i][j] = uint8((i + j)/2)
            */
        }
    }
    return z
}

func main() {
    pic.Show(Pic)
}

