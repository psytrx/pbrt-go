package scenes

import (
	"pbrt/pkg/pbrt"
	"pbrt/pkg/pbrt/vec"
)

func NewSpheres(aspectRatio float64) pbrt.Scene {
	lookFrom := vec.New(0, 2, -6)
	lookAt := vec.New(0, 0.5, 0)
	focusDist := lookAt.Sub(lookFrom).Len()

	white := pbrt.NewLambertian(vec.New(1, 1, 1))
	red := pbrt.NewLambertian(vec.New(1, 0, 0))
	green := pbrt.NewLambertian(vec.New(0, 1, 0))
	blue := pbrt.NewLambertian(vec.New(0, 0, 1))

	scene := pbrt.Scene{
		Camera: pbrt.NewCamera(
			lookFrom, lookAt,
			vec.New(0, -1, 0),
			60,
			aspectRatio,
			0.01, focusDist,
		),
		Background: pbrt.NewLinearGradient(
			vec.New(0.82, 0.55, 0.24),
			vec.New(0.24, 0.45, 0.72),
		),
		World: pbrt.NewList(
			pbrt.NewSphere(vec.New(0, -999, 0), 999, white),
			pbrt.NewSphere(vec.New(-2, 1, 0), 1, red),
			pbrt.NewSphere(vec.New(0, 1, 0), 1, green),
			pbrt.NewSphere(vec.New(2, 1, 0), 1, blue),
		),
	}

	return scene
}
