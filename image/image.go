package image

import (
	"os"
	"raytracing-in-go/colour"
	"strconv"
	"strings"
)

type Image struct {
	pixels        [][]colour.Colour
	Width, Height int
}

func New(width, height int) Image {
	var image Image
	image.Width = width
	image.Height = height

	for _ = range width {
		var newCol = make([]colour.Colour, height)

		image.pixels = append(image.pixels, newCol)
	}

	return image
}

func (i *Image) GetPixel(x, y int) colour.Colour {
	return i.pixels[x][y]
}

func (i *Image) SetPixel(x, y int, colour colour.Colour) {
	i.pixels[x][y] = colour
}

func WriteToFile(image Image, filepath string) {
	file, _ := os.Create(filepath)
	defer file.Close()

	outputString := strings.Builder{}

	outputString.WriteString(("P3\n"))
	outputString.WriteString(strconv.Itoa(image.Width) + " " + strconv.Itoa(image.Height) + "\n")
	outputString.WriteString(strconv.Itoa(255) + "\n")

	for y := range image.Height {
		for x := range image.Width {
			outputString.WriteString(strconv.Itoa(int(image.GetPixel(x, y).R)))
			outputString.WriteString(" ")
			outputString.WriteString(strconv.Itoa(int(image.GetPixel(x, y).G)))
			outputString.WriteString(" ")
			outputString.WriteString(strconv.Itoa(int(image.GetPixel(x, y).B)))
			outputString.WriteString(" ")
		}
		outputString.WriteString("\n")
	}

	file.WriteString(outputString.String())
}
