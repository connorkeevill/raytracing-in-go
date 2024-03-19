package primitive

import "raytracing-in-go/geometry"

type Hittable interface {
	Material() Material
}

type Hit struct {
	T                float64 // The distance along the ray of the hit.
	IncomingRay      *Ray
	Surface          *Hittable
	Normal, Position geometry.Vector
}
