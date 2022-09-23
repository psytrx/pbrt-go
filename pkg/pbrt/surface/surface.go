package surface

import "pbrt/pkg/pbrt/ray"

type Surface interface {
	Intersect(r ray.Ray, tMin, tMax float64) (bool, *Intersection)
}
