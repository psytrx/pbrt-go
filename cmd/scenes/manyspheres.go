package scenes

import (
	"math/rand"

	"pbrt/pkg/pbrt"
	"pbrt/pkg/pbrt/vec"
)

func NewManySpheres(aspectRatio float64) pbrt.Scene {
	rng := rand.New(rand.NewSource(42))

	lookFrom := vec.New(8, 2, -4)
	lookAt := vec.New(0, 0.5, 0)
	focusDist := lookAt.Sub(lookFrom).Len()

	white := pbrt.NewLambertian(vec.New(1, 1, 1))
	brown := pbrt.NewLambertian(vec.New(0.4, 0.2, 0.1))
	glass := pbrt.NewDielectric(vec.New(0.95, 0.95, 1), 1.5)
	metal := pbrt.NewMetal(vec.New(0.2, 0.9, 0.2), 0.05)

	surfaces := []pbrt.Surface{
		pbrt.NewSphere(vec.New(0, -999, 0), 999, white),
		pbrt.NewSphere(vec.New(-2, 1, 0), 1, brown),
		pbrt.NewSphere(vec.New(0, 1, 0), 1, glass),
		pbrt.NewSphere(vec.New(2, 1, 0), 1, metal),
	}

	for i := 0; i < 128; i++ {
		radius := 0.2 + 0.2*rng.Float64()
		disk := vec.RandomInUnitDisk(rng).Scaled(16)
		pos := vec.New(disk.X, radius, disk.Y)

		// inside the big spheres
		if pos.Len() < 3 {
			continue
		}

		rand := rng.Float64()
		var mat pbrt.Material

		if rand < 0.5 {
			clr := vec.RandomUniform(0, 1, rng)
			clrSqr := clr.Mult(clr)
			mat = pbrt.NewLambertian(clrSqr)
		} else if rand < 0.75 {
			clr := vec.RandomUniform(0.5, 1, rng)
			mat = pbrt.NewMetal(clr, 0.25*rng.Float64())
		} else {
			clr := vec.RandomUniform(0.9, 1, rng)
			mat = pbrt.NewDielectric(clr, 1.5)
		}

		sphere := pbrt.NewSphere(pos, radius, mat)
		surfaces = append(surfaces, sphere)
	}

	scene := pbrt.Scene{
		Camera: pbrt.NewCamera(
			lookFrom, lookAt,
			vec.New(0, -1, 0),
			55,
			aspectRatio,
			0.05, focusDist,
		),
		Background: pbrt.NewLinearGradient(
			vec.New(0.82, 0.55, 0.24),
			vec.New(0.24, 0.45, 0.72),
		),
		World: pbrt.NewList(surfaces...),
	}

	return scene
}
