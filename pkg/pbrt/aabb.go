package pbrt

import (
	"math"
	"pbrt/pkg/pbrt/vec"
)

type AABB struct {
	min, max vec.Vec
}

func (aabb AABB) Intersects(r *Ray, tMin, tMax float64) bool {
	t0 := math.Min(
		(aabb.min.X-r.Origin.X)/r.Direction.X,
		(aabb.max.X-r.Origin.X)/r.Direction.X)
	t1 := math.Max(
		(aabb.min.X-r.Origin.X)/r.Direction.X,
		(aabb.max.X-r.Origin.X)/r.Direction.X)

	tMin = math.Max(t0, tMin)
	tMax = math.Min(t1, tMax)

	if tMax <= tMin {
		return false
	}

	t0 = math.Min(
		(aabb.min.Y-r.Origin.Y)/r.Direction.Y,
		(aabb.max.Y-r.Origin.Y)/r.Direction.Y)
	t1 = math.Max(
		(aabb.min.Y-r.Origin.Y)/r.Direction.Y,
		(aabb.max.Y-r.Origin.Y)/r.Direction.Y)

	tMin = math.Max(t0, tMin)
	tMax = math.Min(t1, tMax)

	if tMax <= tMin {
		return false
	}

	t0 = math.Min(
		(aabb.min.Z-r.Origin.Z)/r.Direction.Z,
		(aabb.max.Z-r.Origin.Z)/r.Direction.Z)
	t1 = math.Max(
		(aabb.min.Z-r.Origin.Z)/r.Direction.Z,
		(aabb.max.Z-r.Origin.Z)/r.Direction.Z)

	tMin = math.Max(t0, tMin)
	tMax = math.Min(t1, tMax)

	if tMax <= tMin {
		return false
	}

	return true
}

func (aabb AABB) Area() float64 {
	a := aabb.max.X - aabb.min.X
	b := aabb.max.Y - aabb.min.Z
	c := aabb.max.Z - aabb.min.Z
	return 2 * (a*b + b*c + c*a)
}

func surroundingBox(a, b AABB) AABB {
	small := vec.New(
		math.Min(a.min.X, b.min.X),
		math.Min(a.min.Y, b.min.Y),
		math.Min(a.min.Z, b.min.Z))
	big := vec.New(
		math.Max(a.max.X, b.max.X),
		math.Max(a.max.Y, b.max.Y),
		math.Max(a.max.Z, b.max.Z))
	return AABB{small, big}
}
