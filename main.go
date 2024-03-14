package main

import (
	"raytracing-in-go/colour"
	"raytracing-in-go/image"
)

func main() {
	width := 50
	height := 30

	picture := image.New(width, height)

	for row := range height {
		for col := range width {
			if (row+col)%2 == 0 {
				picture[row][col] = colour.New(255, 0, 0)
			} else {
				picture[row][col] = colour.New(0, 0, 255)
			}
		}
	}

	// By this point we have a simple checkerboard image; now let's save to a file:
	image.WriteToFile(picture, "image.ppm")
}
