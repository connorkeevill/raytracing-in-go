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
