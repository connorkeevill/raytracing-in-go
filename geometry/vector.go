package geometry

import (
	"math"
)

type Vector struct {
	X, Y, Z float64
}

func (this *Vector) Normalise() {
	length := this.Magnitude()

	this.X /= length
	this.Y /= length
	this.Z /= length
}

func (this *Vector) Magnitude() float64 {
	return math.Sqrt(math.Pow(this.X, 2) + math.Pow(this.Y, 2) + math.Pow(this.Z, 2))
}

func (this *Vector) Dot(other *Vector) float64 {
	return this.X*other.X + this.Y*other.Y + this.Z*other.Z
}

func (this *Vector) Cross(other *Vector) Vector {
	return Vector{
		X: this.Y*other.Z - this.Z*other.Y,
		Y: this.Z*other.X - this.X*other.Z,
		Z: this.X*other.Y - this.Y*other.X,
	}
}

func (this *Vector) Add(other *Vector) Vector {
	return Vector{
		X: this.X + other.X,
		Y: this.Y + other.Y,
		Z: this.Z + other.Z,
	}
}

func (this *Vector) Times(constant float64) Vector {
	return Vector{
		X: this.X * constant,
		Y: this.Y * constant,
		Z: this.Z * constant,
	}
}

func VectorSum(vectors ...*Vector) Vector {
	var result Vector

	for _, vector := range vectors {
		result = result.Add(vector)
	}

	return result
}
