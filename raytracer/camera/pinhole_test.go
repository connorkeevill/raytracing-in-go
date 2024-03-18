package camera

import (
	"raytracing-in-go/colour"
	"raytracing-in-go/geometry"
	"raytracing-in-go/raytracer/primitive"
	"testing"
)

// MockTraceable implements the Traceable interface for testing.
type MockTraceable struct{}

func (m *MockTraceable) Trace(ray primitive.Ray) colour.Colour {
	// Simple mock implementation: returns a fixed colour for any ray.
	return &colour.RGB8Bit{R: 1.0, G: 0.0, B: 0.0}
}

func TestNewPinholeCamera(t *testing.T) {
	position := geometry.Vector{X: 1, Y: 2, Z: 3}
	lookDir := geometry.Vector{X: 0, Y: 0, Z: -1}
	resolution := Resolution{Width: 800, Height: 600}
	fov := 90.0

	cam := New(position, lookDir, resolution, fov)

	// Test if the camera is initialized correctly
	if cam.FOV != fov {
		t.Errorf("Expected FOV %v, got %v", fov, cam.FOV)
	}
	if cam.Sensor.Width != resolution.Width || cam.Sensor.Height != resolution.Height {
		t.Errorf("Sensor resolution does not match expected values")
	}
	// Further tests can assert the correctness of `origin`, `lookDir`, `upDir`, `rightDir`, and `planeDistance`
}

func TestGetRayForPixel(t *testing.T) {
	position := geometry.Vector{X: 0, Y: 0, Z: 0}
	lookDir := geometry.Vector{X: 0, Y: 0, Z: -1}
	resolution := Resolution{Width: 800, Height: 600}
	fov := 90.0

	cam := New(position, lookDir, resolution, fov)

	ray := cam.GetRayForPixel(resolution.Width/2, resolution.Height/2)

	if ray.Start != position {
		t.Errorf("Ray start does not match expected position")
	}

	v := geometry.Vector{X: 0, Y: 0, Z: -1}

	if ray.Direction != v {
		t.Errorf("Ray direction does not match expected direction")
	}
	// Test for other pixels as needed
}

func TestRender(t *testing.T) {
	position := geometry.Vector{X: 0, Y: 0, Z: 0}
	lookDir := geometry.Vector{X: 0, Y: 0, Z: -1}
	resolution := Resolution{Width: 100, Height: 100} // Smaller resolution for faster test
	fov := 90.0

	cam := New(position, lookDir, resolution, fov)
	traceable := MockTraceable{}

	rendered := cam.Render(&traceable)

	if rendered.Width != resolution.Width || rendered.Height != resolution.Height {
		t.Errorf("Rendered image resolution does not match expected values")
	}

	// Check if the mock Traceable was called and the correct colour was set for at least one pixel
	expectedColour := colour.RGB8Bit{R: 1.0, G: 0.0, B: 0.0}
	if pixelColour := rendered.GetPixel(0, 0); pixelColour != expectedColour {
		t.Errorf("Expected pixel colour %v, got %v", expectedColour, pixelColour)
	}
	// Additional tests could iterate over all pixels if needed
}
