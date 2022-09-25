package pbrt

import (
	"math"

	"pbrt/pkg/pbrt/vec"
)

type Sphere struct {
	center   vec.Vec
	radius   float64
	material Material
}

func NewSphere(center vec.Vec, radius float64, material Material) Sphere {
	return Sphere{center, radius, material}
}

func (s Sphere) Intersect(r *Ray, tMin, tMax float64) (bool, *Intersection) {
	oc := r.Origin.Sub(s.center)
	a := r.Direction.LenSqr()
	halfB := vec.Dot(oc, r.Direction)
	c := oc.LenSqr() - s.radius*s.radius

	d := halfB*halfB - a*c
	if d < 0 {
		return false, nil
	}
	sqrtD := math.Sqrt(d)

	root := (-halfB - sqrtD) / a
	if root < tMin || root > tMax {
		root = (-halfB + sqrtD) / a
		if root < tMin || root > tMax {
			return false, nil
		}
	}

	t := root
	p := r.At(t)
	outwardNormal := p.Sub(s.center).Scaled(1 / s.radius)
	isect := NewIsect(r, t, p, outwardNormal, s.material)

	return true, &isect
}
