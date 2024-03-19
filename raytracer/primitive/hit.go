package primitive

import "raytracing-in-go/geometry"

type Hit struct {
	T                float64 // The distance along the ray of the hit.
	IncomingRay      *Ray
	Surface          *Object
	Normal, Position geometry.Vector
}
