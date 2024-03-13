package main

import (
	"os"
	"strconv"
)

type colour struct {
	r, g, b byte // We are using 8-bit colour
}

func writeImageToFile(image [][]colour, filepath string) {
	file, _ := os.Create(filepath)
	defer file.Close()

	file.WriteString(("P3\n"))
	file.WriteString(strconv.Itoa(len(image[0])) + " " + strconv.Itoa(len(image)) + "\n")
	file.WriteString(strconv.Itoa(255) + "\n")

	for _, row := range image {
		for _, column := range row {
			file.WriteString(strconv.Itoa(int(column.r)))
			file.WriteString(" ")
			file.WriteString(strconv.Itoa(int(column.g)))
			file.WriteString(" ")
			file.WriteString(strconv.Itoa(int(column.b)))
			file.WriteString(" ")
		}
		file.WriteString("\n")
	}
}

func main() {

	width := 50
	height := 30

	var image [][]colour

	red := colour{255, 0, 0}

	for row := range height {
		var newRow = make([]colour, width)

		for column := range width {
			if (row+column)%2 == 0 {
				newRow[column] = red
			}
		}

		image = append(image, newRow)
	}

	// By this point we have a simple checkerboard image; now let's save to a file:
	writeImageToFile(image, "output.ppm")
}
