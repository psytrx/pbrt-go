package cam

import (
	"math"
	"pbrt/pkg/pbrt/ray"
	"pbrt/pkg/pbrt/vec"
)

type Camera struct {
	origin               vec.Vec
	topLeftCorner        vec.Vec
	horizontal, vertical vec.Vec
	u, v, w              vec.Vec
	lensRadius           float64
}

func New(
	lookFrom, lookAt vec.Vec,
	vDown vec.Vec,
	vFov float64,
	aspectRatio float64,
	aperture, focusDist float64,
) Camera {
	theta := degToRad(vFov)
	h := math.Tan(theta / 2)
	viewportHeight := 2 * h
	viewportWidth := aspectRatio * viewportHeight

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

func (c Camera) Ray(s, t float64) ray.Ray {
	rd := vec.Zero()
	offset := c.u.Scaled(rd.X).Add(c.v.Scaled(rd.Y))
	return ray.New(
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
