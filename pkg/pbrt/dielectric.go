package pbrt

import (
	"math"
	"math/rand"

	"pbrt/pkg/pbrt/vec"
)

type Dielectric struct {
	albedo            vec.Vec
	indexOfRefraction float64
}

func NewDielectric(albedo vec.Vec, indexOfRefraction float64) Dielectric {
	return Dielectric{albedo, indexOfRefraction}
}

func (d Dielectric) Scatter(r *Ray, isect *Intersection, rng *rand.Rand) (bool, *vec.Vec, *Ray) {
	var refractionRatio float64
	if isect.FrontFace {
		refractionRatio = 1.0 / d.indexOfRefraction
	} else {
		refractionRatio = d.indexOfRefraction
	}

	cosTheta := math.Min(1, vec.Dot(r.Direction.Scaled(-1), isect.Normal))
	sinTheta := math.Sqrt(1 - cosTheta*cosTheta)

	cannotRefract := refractionRatio*sinTheta > 1

	var direction vec.Vec
	if cannotRefract || reflectance(cosTheta, refractionRatio) > rng.Float64() {
		direction = vec.Reflect(r.Direction, isect.Normal)
	} else {
		direction = vec.Refract(r.Direction, isect.Normal, refractionRatio)
	}

	scattered := NewRay(isect.P, direction)
	return true, &d.albedo, &scattered
}

func reflectance(cosine, refIdx float64) float64 {
	r0 := (1.0 - refIdx) / (1.0 + refIdx)
	r0 *= r0
	return r0 + (1.0-r0)*math.Pow(1-cosine, 5)
}
