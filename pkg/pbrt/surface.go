package pbrt

type Surface interface {
	Intersect(r *Ray, tMin, tMax float64) (bool, *Intersection)
}
