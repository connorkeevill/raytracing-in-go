package object

import "raytracing-in-go/raytracer/primitive"

type Mesh struct {
}

func (mesh *Mesh) Intersect(ray *primitive.Ray) []primitive.Hit {
	return nil
}

func (mesh *Mesh) Material() primitive.Material {
	return nil
}
