package camera

import (
	"raytracing-in-go/colour"
	"raytracing-in-go/raytracer/primitive"
)

type DemoTraceable struct{}

func (env *DemoTraceable) Trace(ray *primitive.Ray) colour.Colour {
	return colour.RGB8Bit{
		R: byte((ray.Direction.X + 1) * 128),
		G: byte((ray.Direction.Y + 1) * 128),
		B: byte((ray.Direction.Z + 1) * 128),
	}
}
