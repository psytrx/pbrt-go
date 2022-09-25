package pbrt

import (
	"math/rand"

	"pbrt/pkg/pbrt/vec"
)

type Lambertian struct {
	albedo vec.Vec
}

func NewLambertian(albedo vec.Vec) Lambertian {
	return Lambertian{albedo}
}

func (m Lambertian) Scatter(r Ray, isect Intersection, rng *rand.Rand) (bool, *vec.Vec, *Ray) {
	direction := vec.RandomInHemisphere(isect.Normal, rng)
	scattered := NewRay(isect.P, direction)
	return true, &m.albedo, &scattered
}
