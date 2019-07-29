package rndimage

import (
	"image"
	"image/color"
	"math/rand"
	"sync"
)

// Generate - generates `image.RGBA` of given `width` and `height` filling it with random pixel colors
func Generate(width, height int) *image.RGBA {
	rgba := image.NewRGBA(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{width, height}})

	var wg sync.WaitGroup
	wg.Add(width)

	for i := 0; i < width; i++ {
		go func(i int) {
			rand := rand.New(rand.NewSource(rand.Int63()))
			for k := 0; k < height; k++ {
				rgba.Set(i, k, color.RGBA{
					R: uint8(rand.Intn(256)),
					G: uint8(rand.Intn(256)),
					B: uint8(rand.Intn(256)),
					A: uint8(rand.Intn(256)),
				})
			}
			wg.Done()
		}(i)
	}

	wg.Wait()

	return rgba
}
