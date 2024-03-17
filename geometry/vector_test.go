package geometry

import (
	"testing"
)

// Thanks to ChatGPT for making a start on these unit tests. As this project evolves I will return to implement more
// thorough tests

func TestMagnitude(t *testing.T) {
	v := Vector{X: 3, Y: 4, Z: 0}
	expected := 5.0 // Based on the Pythagorean theorem
	if mag := v.Magnitude(); mag != expected {
		t.Errorf("Expected magnitude of %v to be %v, got %v", v, expected, mag)
	}
}

func TestNormalise(t *testing.T) {
	v := Vector{X: 3, Y: 0, Z: 4}
	v.Normalise()
	expectedMagnitude := 1.0
	if mag := v.Magnitude(); mag != expectedMagnitude {
		t.Errorf("Expected magnitude after normalisation to be %v, got %v", expectedMagnitude, mag)
	}
}

func TestDot(t *testing.T) {
	v1 := Vector{X: 1, Y: 2, Z: 3}
	v2 := Vector{X: 4, Y: 5, Z: 6}
	expected := 32.0 // 1*4 + 2*5 + 3*6
	if dot := v1.Dot(&v2); dot != expected {
		t.Errorf("Expected dot product to be %v, got %v", expected, dot)
	}
}

func TestCross(t *testing.T) {
	v1 := Vector{X: 1, Y: 2, Z: 3}
	v2 := Vector{X: 4, Y: 5, Z: 6}
	expected := Vector{X: -3, Y: 6, Z: -3}
	if cross := v1.Cross(&v2); cross != expected {
		t.Errorf("Expected cross product to be %v, got %v", expected, cross)
	}
}

func TestAdd(t *testing.T) {
	v1 := Vector{X: 1, Y: 2, Z: 3}
	v2 := Vector{X: 4, Y: 5, Z: 6}
	expected := Vector{X: 5, Y: 7, Z: 9}
	if sum := v1.Add(&v2); sum != expected {
		t.Errorf("Expected sum to be %v, got %v", expected, sum)
	}
}

func TestTimes(t *testing.T) {
	v := Vector{X: 1, Y: 2, Z: 3}
	constant := 2.0
	expected := Vector{X: 2, Y: 4, Z: 6}
	if product := v.Times(constant); product != expected {
		t.Errorf("Expected product to be %v, got %v", expected, product)
	}
}

func TestVectorSum(t *testing.T) {
	v1 := &Vector{X: 1, Y: 2, Z: 3}
	v2 := &Vector{X: 4, Y: 5, Z: 6}
	expected := Vector{X: 5, Y: 7, Z: 9}
	if sum := VectorSum(v1, v2); sum != expected {
		t.Errorf("Expected sum to be %v, got %v", expected, sum)
	}
}
