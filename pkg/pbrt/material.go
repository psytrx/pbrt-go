package pbrt

import (
	"math/rand"

	"pbrt/pkg/pbrt/vec"
)

type Material interface {
	Scatter(r *Ray, isect *Intersection, rng *rand.Rand) (bool, *vec.Vec, *Ray)
}
