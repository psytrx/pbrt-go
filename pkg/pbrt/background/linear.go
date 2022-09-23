package background

import (
	"pbrt/pkg/pbrt/ray"
	"pbrt/pkg/pbrt/vec"
)

type LinearGradientBackground struct {
	bottom, top vec.Vec
}

func NewLinearGradient(bottom, top vec.Vec) LinearGradientBackground {
	return LinearGradientBackground{bottom, top}
}

func (bg LinearGradientBackground) RayColor(r ray.Ray) vec.Vec {
	unitDirection := r.Direction.Normalized()
	t := (unitDirection.Y + 1) / 2
	return bg.bottom.Scaled(1 - t).Add(bg.top.Scaled(t))
}
