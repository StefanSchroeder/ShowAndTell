package main

import "math/cmplx"
import "image"
import "image/color"
import "image/png"
import "os"

const limit = 4.0
const iterations = 100
const width = 512.0
const height = 512.0

var rotPixel color.Color = color.RGBA{255, 0, 0, 255}
var greenPixel color.Color = color.RGBA{0, 255, 0, 255}
func main() {

	m := image.NewRGBA(image.Rect(0, 0, width, height))
	for y:=0.0; y<height; y++ {
		for x:=0.0; x<width; x++ {
			m.Set(int(x), int(y), inSet(x,y))
		}
	}
	file, _ := os.Create("mandel.png")
	png.Encode(file, m)
	file.Close()

}

func inSet(x, y float64) color.Color {
	z := complex(0.0, 0.0)
	c := complex(2*x/width - 1.5, 2*y/height - 1.0)

	for i:=0 ; i<iterations ; i++ {
		z = z*z + c
	}
	if cmplx.Abs(z) < limit {
		return rotPixel
	} else {
		return greenPixel
	}
}
