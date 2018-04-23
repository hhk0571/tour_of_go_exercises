/* https://tour.golang.org/methods/20

Sqrt should return a non-nil error value when given a negative number, as it doesn't support complex numbers.

*/

package main

import (
	"fmt"
)

const EPSILON = 0.000000000000001

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0
	old_z := 0.0
	fmt.Println("========BEGIN==========")
	for i := 1; (old_z-z)*(old_z-z) > EPSILON; i++ {
		old_z = z
		z -= (z*z - x) / (2 * z)
		fmt.Println(i, ":", z)
	}
	fmt.Println("=========END===========")

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

