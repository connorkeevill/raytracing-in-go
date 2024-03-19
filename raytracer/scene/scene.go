package scene

import (
	"raytracing-in-go/colour"
	. "raytracing-in-go/raytracer/primitive"
)

type Scene struct {
	Objects []Object
}

func (scene *Scene) Trace(ray *Ray) colour.Colour {
	var hits []Hit

	for _, object := range scene.Objects {
		hits = append(hits, object.Intersect(ray)...)
	}

	if len(hits) == 0 {
		col := colour.NewRGB(0, 0, 0)
		return &col
	}

	// Get the closest hit
	closestHit := &hits[0]

	for _, hit := range hits {
		if hit.T < closestHit.T {
			closestHit = &hit
		}
	}

	position := ray.At(closestHit.T)
	return closestHit.Surface.Material().Shade(closestHit, scene, &position, &closestHit.Normal)
}
