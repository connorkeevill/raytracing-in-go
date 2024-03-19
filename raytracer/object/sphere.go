package object

import (
	"math"
	"raytracing-in-go/geometry"
	"raytracing-in-go/raytracer/primitive"
)

type Sphere struct {
	M      primitive.Material
	Center geometry.Vector
	Radius float64
}

func (sphere *Sphere) Intersect(ray *primitive.Ray) []primitive.Hit {
	sphereCenterToRayStart := ray.Start.Subtract(&sphere.Center)

	a := ray.Direction.Dot(&ray.Direction)
	b := 2 * ray.Direction.Dot(&sphereCenterToRayStart)
	c := sphereCenterToRayStart.Dot(&sphereCenterToRayStart) - (sphere.Radius * sphere.Radius)

	discriminant := (b * b) - 4*a*c

	if discriminant < 0 {
		return []primitive.Hit{}
	}

	var hits []primitive.Hit
	var surface primitive.Object = sphere
	squareRootOf := math.Sqrt(discriminant)

	t := -b + squareRootOf/2*a
	position := ray.At(t)
	hits = append(hits, primitive.Hit{
		T:           t,
		IncomingRay: ray,
		Surface:     &surface,
		Normal:      position.Subtract(&sphere.Center),
		Position:    position,
	})

	if discriminant == 0 {
		return hits
	}

	t = -b - squareRootOf/2*a
	position = ray.At(t)
	hits = append(hits, primitive.Hit{
		T:           t,
		IncomingRay: ray,
		Surface:     &surface,
		Normal:      position.Subtract(&sphere.Center),
		Position:    position,
	})

	return hits
}

func (sphere *Sphere) Material() primitive.Material { return sphere.M }
