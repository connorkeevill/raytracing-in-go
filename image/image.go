package image

import (
	"os"
	"raytracing-in-go/colour"
	"strconv"
)

type Image [][]colour.Colour

func New(width, height int) Image {
	var image [][]colour.Colour

	for _ = range height {
		var newRow = make([]colour.Colour, width)

		image = append(image, newRow)
	}

	return image
}

func WriteToFile(image Image, filepath string) {
	file, _ := os.Create(filepath)
	defer file.Close()

	file.WriteString(("P3\n"))
	file.WriteString(strconv.Itoa(len(image[0])) + " " + strconv.Itoa(len(image)) + "\n")
	file.WriteString(strconv.Itoa(255) + "\n")

	for _, row := range image {
		for _, column := range row {
			file.WriteString(strconv.Itoa(int(column.R)))
			file.WriteString(" ")
			file.WriteString(strconv.Itoa(int(column.G)))
			file.WriteString(" ")
			file.WriteString(strconv.Itoa(int(column.B)))
			file.WriteString(" ")
		}
		file.WriteString("\n")
	}
}
