package ray

import "pbrt/pkg/pbrt/vec"

type Ray struct {
	Origin, Direction vec.Vec
}

func New(origin, direction vec.Vec) Ray {
	return Ray{origin, direction}
}

func (ray Ray) At(t float64) vec.Vec {
	return ray.Origin.Add(ray.Direction.Scaled(t))
}
