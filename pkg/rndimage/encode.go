package rndimage

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
)

// Encode - encodes provided `image.Image` as jpeg
func Encode(writer io.Writer, image image.Image) error {
	err := jpeg.Encode(writer, image, nil)
	if err != nil {
		return fmt.Errorf("unable to encode jpeg: %v", err)
	}
	return nil
}
