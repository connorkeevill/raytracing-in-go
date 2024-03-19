package primitive

import (
	. "raytracing-in-go/geometry"
)

type Ray struct {
	Direction, Start Vector
}

func (ray *Ray) Normalise() {
	ray.Direction.Normalise()
}

func (ray *Ray) At(T float64) Vector {
	distance := ray.Direction.Times(T)

	return ray.Start.Add(&distance)
}
