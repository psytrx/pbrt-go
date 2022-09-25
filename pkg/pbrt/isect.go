package pbrt

import (
	"pbrt/pkg/pbrt/vec"
)

type Intersection struct {
	T        float64
	P        vec.Vec
	Normal   vec.Vec
	Material Material
}

func NewIsect(r Ray, t float64, p, outwardNormal vec.Vec) Intersection {
	frontFace := vec.Dot(r.Direction, outwardNormal) < 0
	if !frontFace {
		outwardNormal = outwardNormal.Scaled(-1)
	}

	return Intersection{
		T:      t,
		P:      p,
		Normal: outwardNormal,
	}
}
