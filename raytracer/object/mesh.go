package object

import (
	. "raytracing-in-go/geometry"
	. "raytracing-in-go/raytracer/primitive"
)

type face struct {
	a, b, c *Vector
	mesh    *Mesh
}

func (f *face) Intersect(ray *Ray) []Hit {
	AB := f.b.Subtract(f.a)
	AC := f.c.Subtract(f.a)
	normal := AB.Cross(&AC)
	normal.Normalise()

	// Ray and plane do not intersect if they are parallel (so we check if ray and normal are perp.)
	normalDotDirection := normal.Dot(&ray.Direction)
	if normalDotDirection == 0 {
		return []Hit{}
	}

	// Plane equation now becomes ax + by + cz - d = 0 where (a b c) = normal and d = [the below]
	d := normal.X*f.a.X + normal.Y*f.a.Y + normal.Z*f.a.Z

	// Solving the above equation gets us this expression for T:
	T := -(normal.Dot(&ray.Start) * d) / normalDotDirection
	position := ray.At(T)

	// The ray intersects with the plane - great. Now must check that the intersection is within the triangle.
	// For this we will use the barycentric coordinates approach.
	u, v, w := f.barycentric(position, AB, AC)

	// Coordinates u, v and w are the Barycentric coordinates. These tell us whether the intersection is within triangle
	if !((0 <= u && u <= 1) && (0 <= v && v <= 1) && (0 <= w && w <= 1)) {
		return []Hit{}
	}

	return []Hit{
		{
			T:           T,
			IncomingRay: *ray,
			Normal:      normal,
			Surface:     f.mesh, //  Note that we are providing a pointer here to avoid copying the entire mesh for each hit.
			Position:    position,
		},
	}
}

func (f *face) barycentric(planeIntersection, AB, AC Vector) (float64, float64, float64) {
	AP := planeIntersection.Subtract(f.a)

	// Some ugly matrix work now to essentially solve a polynomial with Cramers rule.
	d00 := AB.Dot(&AB)
	d01 := AB.Dot(&AC)
	d11 := AC.Dot(&AC)
	d20 := AP.Dot(&AB)
	d21 := AP.Dot(&AC)
	D := d00*d11 - d01*d01

	u := (d11*d20 - d01*d21) / D
	v := (d00*d21 - d01*d20) / D
	w := 1 - u - v
	return u, v, w
}

type Mesh struct {
	vertices map[Vector]struct{}
	faces    []face
	material Material
}

func (mesh *Mesh) Intersect(ray *Ray) []Hit {
	var hits []Hit

	for _, f := range mesh.faces {
		faceIntersection := f.Intersect(ray)

		if len(faceIntersection) == 0 {
			continue
		}

		hits = append(hits, faceIntersection...)
	}

	return hits
}

func (mesh *Mesh) Material() Material {
	return mesh.material
}
