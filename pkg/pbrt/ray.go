package pbrt

import "pbrt/pkg/pbrt/vec"

type Ray struct {
	Origin, Direction vec.Vec
}

func NewRay(origin, direction vec.Vec) Ray {
	return Ray{origin, direction.Normalized()}
}

func (ray Ray) At(t float64) vec.Vec {
	return ray.Origin.Add(ray.Direction.Scaled(t))
}
