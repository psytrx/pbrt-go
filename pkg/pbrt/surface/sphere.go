package surface

import (
	"math"
	"pbrt/pkg/pbrt/ray"
	"pbrt/pkg/pbrt/vec"
)

type Sphere struct {
	center vec.Vec
	radius float64
}

func NewSphere(center vec.Vec, radius float64) Sphere {
	return Sphere{center, radius}
}

func (s Sphere) Intersect(r ray.Ray, tMin, tMax float64) (bool, *Intersection) {
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
	isect := NewIntersection(r, t, p, outwardNormal)

	return true, &isect
}
