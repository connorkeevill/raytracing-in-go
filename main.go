package main

import (
	"fmt"
	"os"
	"raytracing-in-go/geometry"
	"raytracing-in-go/image"
	"raytracing-in-go/raytracer/camera"
	"raytracing-in-go/raytracer/object"
	"raytracing-in-go/raytracer/primitive"
	"raytracing-in-go/raytracer/scene"
)

func main() {
	demo := camera.DemoTraceable{}

	environment := scene.Scene{}
	environment.Objects = append(environment.Objects, &object.Sphere{
		M:      &primitive.RayDirectionMaterial{},
		Center: geometry.Vector{Z: 2.5},
		Radius: 1})

	file, _ := os.Open("teapot.obj")
	mesh := object.FromObjFile(*file)
	mat := &primitive.RayDirectionMaterial{}
	mesh.SetMaterial(mat)
	environment.Objects = append(environment.Objects, &mesh)

	cam := camera.New(geometry.Vector{Y: 0, Z: -35}, geometry.Vector{Z: 1}, camera.Resolution{Width: 500, Height: 500}, 60)
	demoImage := cam.Render(&demo)
	image.WriteToFile(demoImage, "demo.ppm")
	picture := cam.Render(&environment)

	fmt.Println("Rendered, writing to file")
	image.WriteToFile(picture, "image.ppm")
	fmt.Println("Written")
}
