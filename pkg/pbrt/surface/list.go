package surface

import "pbrt/pkg/pbrt/ray"

type SurfaceList []Surface

func NewList(surfaces ...Surface) SurfaceList {
	return SurfaceList(surfaces)
}

func (xs SurfaceList) Intersect(r ray.Ray, tMin, tMax float64) (bool, *Intersection) {
	hitAnything := false
	closestT := tMax
	var closest *Intersection

	for _, s := range xs {
		if ok, isect := s.Intersect(r, tMin, closestT); ok {
			hitAnything = true
			closestT = isect.T
			closest = isect
		}
	}

	return hitAnything, closest
}
