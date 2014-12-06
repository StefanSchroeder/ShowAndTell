package main

import "fmt"
import "math/cmplx"

const limit = 4.0
const iterations = 100
const width = 100.0
const height = 50.0

func main() {

	result := ""
	for y:=0.0; y<height; y++ {
		for x:=0.0; x<width; x++ {
			if inSet(x,y) {
				result += "●"
			} else {
				result += " "
			}
		}
		result += "\n"
	}
	fmt.Printf(result)
}

func inSet(x, y float64) bool {
	z := complex(0.0, 0.0)
	c := complex(2*x/width - 1.5, 2*y/height - 1.0)

	for i:=0 ; i<iterations ; i++ {
		z = z*z + c
	}
	if cmplx.Abs(z) < limit {
		return true
	} else {
		return false
	}
}
