package image

import (
	"os"
	"raytracing-in-go/colour"
	"testing"
)

// Thanks to ChatGPT for these unit tests. More tests will be added as the project evolves

func TestNew(t *testing.T) {
	width, height := 10, 10
	img := New(width, height)

	if img.Width != width || img.Height != height {
		t.Errorf("Expected image dimensions to be %dx%d, but got %dx%d", width, height, img.Width, img.Height)
	}

	if len(img.pixels) != width {
		t.Errorf("Expected pixel width to be %d, but got %d", width, len(img.pixels))
	}

	for _, col := range img.pixels {
		if len(col) != height {
			t.Errorf("Expected pixel height to be %d, but got %d", height, len(col))
		}
	}
}

func TestGetSetPixel(t *testing.T) {
	img := New(10, 10)
	x, y := 5, 5
	expectedColour := colour.Colour{R: 255, G: 255, B: 255}
	img.SetPixel(x, y, expectedColour)

	receivedColour := img.GetPixel(x, y)
	if receivedColour != expectedColour {
		t.Errorf("Expected pixel colour to be %v, but got %v", expectedColour, receivedColour)
	}
}

func TestWriteToFile(t *testing.T) {
	img := New(2, 2)
	colour1 := colour.Colour{R: 255, G: 0, B: 0}
	colour2 := colour.Colour{R: 0, G: 255, B: 0}
	colour3 := colour.Colour{R: 0, G: 0, B: 255}
	colour4 := colour.Colour{R: 255, G: 255, B: 255}
	img.SetPixel(0, 0, colour1)
	img.SetPixel(1, 0, colour2)
	img.SetPixel(0, 1, colour3)
	img.SetPixel(1, 1, colour4)

	filepath := "test_image.ppm"
	WriteToFile(img, filepath)

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		t.Errorf("File %s was not created", filepath)
	}

	// Further tests could check the file content matches expected PPM format
	// Cleanup
	os.Remove(filepath)
}

func TestWriteToFileWithFormatCheck(t *testing.T) {
	img := New(2, 2)
	colour1 := colour.Colour{R: 255, G: 0, B: 0}     // Red
	colour2 := colour.Colour{R: 0, G: 255, B: 0}     // Green
	colour3 := colour.Colour{R: 0, G: 0, B: 255}     // Blue
	colour4 := colour.Colour{R: 255, G: 255, B: 255} // White
	img.SetPixel(0, 0, colour1)
	img.SetPixel(1, 0, colour2)
	img.SetPixel(0, 1, colour3)
	img.SetPixel(1, 1, colour4)

	filepath := "test_image.ppm"
	WriteToFile(img, filepath)
	defer os.Remove(filepath) // Cleanup after the test

	file, err := os.ReadFile(filepath)
	if err != nil {
		t.Fatalf("Failed to read file %s: %v", filepath, err)
	}

	fileContent := string(file)
	expectedContent := "P3\n2 2\n255\n255 0 0 0 255 0 \n0 0 255 255 255 255 \n"

	if fileContent != expectedContent {
		t.Errorf("File content did not match expected format.\nExpected:\n%s\nGot:\n%s", expectedContent, fileContent)
	}
}
