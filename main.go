package main

import (
	"fmt"
	"raytracing-in-go/geometry"
	"raytracing-in-go/image"
	"raytracing-in-go/raytracer/camera"
)

func main() {
	cam := camera.New(geometry.Vector{}, geometry.Vector{Z: 1}, camera.Resolution{Width: 5000, Height: 5000}, 170)

	tracingEnv := camera.DemoTraceable{}
	picture := cam.Render(&tracingEnv)

	fmt.Println("Rendered, writing to file")

	image.WriteToFile(picture, "image.ppm")

	fmt.Println("Written")
}
