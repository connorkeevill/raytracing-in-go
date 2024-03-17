package camera

import (
	"math"
	"raytracing-in-go/colour"
	"raytracing-in-go/geometry"
	"raytracing-in-go/image"
	"raytracing-in-go/raytracer/primitive"
)

type Pinhole struct {
	FOV           float64
	Sensor        image.Image
	origin        geometry.Vector
	lookDir       geometry.Vector
	upDir         geometry.Vector
	rightDir      geometry.Vector
	planeDistance float64
}

type Resolution struct {
	Width, Height int
}

type Traceable interface {
	Trace(ray primitive.Ray) colour.Colour
}

func New(position, lookDir geometry.Vector, resolution Resolution, fov float64) Pinhole {
	// Some ugly looking linear alg coming here; this forms the camera's basis vectors.
	// Firstly: we will work out the vector that points to the right of the camera. To do this, we need to take a dot
	// product between the lookDir vector (or, strictly speaking, its negation), and one other vector; we will use
	// the world "up" vector, assuming that the camera will rotate around the y-axis.
	worldUp := geometry.Vector{Y: 1}
	cameraRight := lookDir.Cross(&worldUp)

	// Now that we have the both the directions in front of and to the right of the camera (i.e., directions in world
	// space) we can take another cross product to get the perpendicular direction; the vector up relative to the camera.
	cameraUp := cameraRight.Cross(&lookDir)

	return Pinhole{
		FOV:      fov,
		Sensor:   image.New(resolution.Width, resolution.Height),
		origin:   position,
		lookDir:  lookDir,
		upDir:    cameraUp,
		rightDir: cameraRight,
		// Maybe a slightly weird expression, but imagine the triangle formed by the camera frustum (from above), and
		// do some trig to work out the distance between the camera center and the opposite side (of the triangle).
		planeDistance: float64(resolution.Width/2) * math.Tan(math.Pi/180*(180-(fov/2))),
	}
}

func (camera *Pinhole) Render(traceable Traceable) image.Image {
	for x := range camera.Sensor.Width {
		for y := range camera.Sensor.Height {
			// Loop variables captured by function may have unintended values (and IDE error??)
			x := x
			y := y
			go func() { camera.Sensor.SetPixel(x, y, traceable.Trace(camera.GetRayForPixel(x, y))) }()
		}
	}

	return camera.Sensor
}

func (camera *Pinhole) GetRayForPixel(x, y int) primitive.Ray {
	var ray primitive.Ray

	// All rays start from the camera center
	ray.Start = camera.origin

	// The direction of the ray will be a linear combination of the camera's basis vectors
	forwardComponent := camera.lookDir.Times(camera.planeDistance)
	rightComponent := camera.rightDir.Times(float64(x - (camera.Sensor.Width / 2)))
	upComponent := camera.upDir.Times(float64(y - (camera.Sensor.Height / 2)))
	ray.Direction = geometry.VectorSum(&forwardComponent, &rightComponent, &upComponent)

	ray.Normalise()

	return ray
}
