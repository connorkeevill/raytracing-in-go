package primitive

import "raytracing-in-go/geometry"

type Hittable interface {
	GetMaterial() Material
}

type Hit struct {
	IncomingRay *Ray
	What        *Hittable
	Normal      geometry.Vector
}
