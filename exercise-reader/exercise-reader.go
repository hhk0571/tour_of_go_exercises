/* https://tour.golang.org/methods/22

Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
*/

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(buf []byte) (int, error) {
	for i := range buf {
		buf[i] = 65
	}
	return len(buf), nil
}

func main() {
	reader.Validate(MyReader{})
}



