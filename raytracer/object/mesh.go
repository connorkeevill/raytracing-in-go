package object

import (
	"bufio"
	"fmt"
	"os"
	. "raytracing-in-go/geometry"
	. "raytracing-in-go/raytracer/primitive"
	"strconv"
	"strings"
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
	T := -(normal.Dot(&ray.Start) + d) / normalDotDirection
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
	vertices map[Vector]*Vector
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

func valueFromErrorTuple[T any](value T, err error) T {
	if err != nil {
		panic("An error occurred")
	}

	return value
}

func (mesh *Mesh) SetMaterial(material Material) {
	mesh.material = material
}

func FromObjFile(file os.File) Mesh {
	fmt.Println("reading mesh")
	mesh := Mesh{
		vertices: map[Vector]*Vector{},
		faces:    []face{},
		material: &RayDirectionMaterial{},
	}

	scanner := bufio.NewScanner(&file)
	var vectors []Vector

	for scanner.Scan() {
		tokens := strings.Fields(scanner.Text())

		if len(tokens) == 0 {
			continue
		}

		if tokens[0] != "v" && tokens[0] != "f" {
			continue
		}

		line := 1

		if tokens[0] == "v" { // Add a new vertex
			newVertex := Vector{
				X: valueFromErrorTuple(strconv.ParseFloat(tokens[1], 64)),
				Y: valueFromErrorTuple(strconv.ParseFloat(tokens[2], 64)),
				Z: valueFromErrorTuple(strconv.ParseFloat(tokens[3], 64)),
			}

			_, ok := mesh.vertices[newVertex]

			// Only add if new
			if !ok {
				mesh.vertices[newVertex] = &newVertex
				vectors = append(vectors, newVertex)
				line += 1
			}
		} else if tokens[0] == "f" {
			// .obj files which contain textures and normals will have multiple indices separated by slashes
			for index := range tokens {
				tokens[index] = strings.Split(tokens[index], "/")[0]
			}

			firstVertex := vectors[valueFromErrorTuple(strconv.Atoi(tokens[1]))-1]

			for index := 2; index < len(tokens); index += 1 {
				mesh.faces = append(mesh.faces, face{
					a:    mesh.vertices[firstVertex],
					b:    mesh.vertices[vectors[valueFromErrorTuple(strconv.Atoi(tokens[index]))-1]],
					c:    mesh.vertices[vectors[valueFromErrorTuple(strconv.Atoi(tokens[index-1]))-1]],
					mesh: &mesh,
				})
			}
		}
	}

	return mesh
}
