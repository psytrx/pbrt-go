package pbrt

import (
	"math/rand"
	"pbrt/pkg/pbrt/vec"
)

type Metal struct {
	albedo vec.Vec
	fuzz   float64
}

func NewMetal(albedo vec.Vec, fuzz float64) Metal {
	return Metal{albedo, fuzz}
}

func (m Metal) Scatter(r *Ray, isect *Intersection, rng *rand.Rand) (bool, *vec.Vec, *Ray) {
	reflected := vec.Reflect(r.Direction, isect.Normal)
	direction := reflected.Add(vec.RandomInUnitSphere(rng))
	scattered := NewRay(isect.P, direction)
	return vec.Dot(scattered.Direction, isect.Normal) > 0, &m.albedo, &scattered
}
