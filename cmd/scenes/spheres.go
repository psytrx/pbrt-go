package scenes

import (
	"pbrt/pkg/pbrt"
	"pbrt/pkg/pbrt/background"
	"pbrt/pkg/pbrt/camera"
	"pbrt/pkg/pbrt/surface"
	"pbrt/pkg/pbrt/vec"
)

func NewSpheres(aspectRatio float64) pbrt.Scene {
	lookFrom := vec.New(0, 2, -6)
	lookAt := vec.New(0, 0.5, 0)
	focusDist := lookAt.Sub(lookFrom).Len()

	scene := pbrt.Scene{
		Camera: camera.New(
			lookFrom, lookAt,
			vec.New(0, -1, 0),
			60,
			aspectRatio,
			0.01, focusDist,
		),
		Background: background.NewLinearGradient(
			vec.New(0.82, 0.55, 0.24),
			vec.New(0.24, 0.45, 0.72),
		),
		World: surface.NewList(
			surface.NewSphere(vec.New(0, 1, 0), 1),
			surface.NewSphere(vec.New(-2, 1, 0), 1),
			surface.NewSphere(vec.New(2, 1, 0), 1),
			surface.NewSphere(vec.New(0, -999, 0), 999),
		),
	}

	return scene
}
