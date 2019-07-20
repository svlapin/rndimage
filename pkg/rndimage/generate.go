package rndimage

import (
	"image"
	"image/color"
	"math/rand"
)

func randomUint8() uint8 {
	return uint8(rand.Intn(256))
}

// Generate - generates `image.RGBA` of given `width` and `height` filling it with random pixel colors
func Generate(width, height int) *image.RGBA {
	rgba := image.NewRGBA(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{width, height}})

	for i := 0; i < width; i++ {
		for k := 0; k < height; k++ {
			rgba.Set(i, k, color.RGBA{
				R: randomUint8(),
				G: randomUint8(),
				B: randomUint8(),
				A: randomUint8(),
			})
		}
	}

	return rgba
}
