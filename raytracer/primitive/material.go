package primitive

import (
	"raytracing-in-go/colour"
	"raytracing-in-go/geometry"
)

type Traceable interface {
	Trace(ray *Ray) colour.Colour
}

type Material interface {
	Shade(
		hit *Hit,
		scene *Traceable,
		position, normal *geometry.Vector) colour.Colour
}

type RayDirectionMaterial struct{}

func (mat *RayDirectionMaterial) Shade(hit *Hit, scene *Traceable, position, normal *geometry.Vector) colour.Colour {
	// Hit normal should already be normalised, but ensure:
	hit.Normal.Normalise()

	matColour := colour.New8Bit(
		byte((hit.Normal.X+1)*128),
		byte((hit.Normal.Y+1)*128),
		byte((hit.Normal.Z+1)*128),
	)

	return &matColour
}
