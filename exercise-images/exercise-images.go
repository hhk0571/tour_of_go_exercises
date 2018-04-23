/* https://tour.golang.org/methods/25

Define your own Image type, implement the necessary methods, and call pic.ShowImage.
*/

package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	w, h  int
	color uint8
}

func (r *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (r *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, r.w, r.h)
}

func (r *Image) At(x, y int) color.Color {
	return color.RGBA{r.color + uint8(x), r.color + uint8(y), 255, 255}
}

func main() {
	m := Image{200, 200, 150}
	pic.ShowImage(&m)
}

