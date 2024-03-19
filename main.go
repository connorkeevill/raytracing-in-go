package main

import (
	"fmt"
	"raytracing-in-go/geometry"
	"raytracing-in-go/image"
	"raytracing-in-go/raytracer/camera"
	"raytracing-in-go/raytracer/object"
	"raytracing-in-go/raytracer/primitive"
	"raytracing-in-go/raytracer/scene"
)

func main() {
	environment := scene.Scene{}
	environment.Objects = append(environment.Objects, &object.Sphere{
		M:      &primitive.RayDirectionMaterial{},
		Center: geometry.Vector{Z: 3},
		Radius: 2})

	cam := camera.New(geometry.Vector{}, geometry.Vector{Z: 1}, camera.Resolution{Width: 1000, Height: 1000}, 90)
	picture := cam.Render(&environment)

	fmt.Println("Rendered, writing to file")
	image.WriteToFile(picture, "image.ppm")
	fmt.Println("Written")
}
