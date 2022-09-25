package pbrt

import (
	"math"
	"math/rand"

	"pbrt/pkg/pbrt/vec"
)

type Camera struct {
	origin               vec.Vec
	topLeftCorner        vec.Vec
	horizontal, vertical vec.Vec
	u, v, w              vec.Vec
	lensRadius           float64
}

func NewCamera(
	lookFrom, lookAt vec.Vec,
	vDown vec.Vec,
	fov float64,
	aspectRatio float64,
	aperture, focusDist float64,
) Camera {
	theta := degToRad(fov)
	halfWidth := math.Tan(theta / 2)
	viewportWidth := 2 * halfWidth
	viewportHeight := viewportWidth / aspectRatio

	w := lookFrom.Sub(lookAt).Normalized()
	u := vec.Cross(vDown, w).Normalized()
	v := vec.Cross(w, u)

	origin := lookFrom
	horizontal := u.Scaled(focusDist * viewportWidth)
	vertical := v.Scaled(focusDist * viewportHeight)
	topLeftCorner := origin.
		Sub(horizontal.Scaled(0.5)).
		Sub(vertical.Scaled(0.5)).
		Sub(w.Scaled(focusDist))

	lensRadius := aperture / 2

	return Camera{
		origin,
		topLeftCorner,
		horizontal, vertical,
		u, v, w,
		lensRadius,
	}
}

func (c Camera) Ray(s, t float64, rng *rand.Rand) Ray {
	rd := vec.RandomInUnitDisk(rng).Scaled(c.lensRadius)
	offset := c.u.Scaled(rd.X).Add(c.v.Scaled(rd.Y))
	return NewRay(
		c.origin,
		c.topLeftCorner.
			Add(c.horizontal.Scaled(s)).
			Add(c.vertical.Scaled(t)).
			Sub(c.origin).
			Sub(offset),
	)
}

func degToRad(deg float64) float64 {
	return deg * math.Pi / 180
}
