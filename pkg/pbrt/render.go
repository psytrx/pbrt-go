package pbrt

import (
	"math"
	"math/rand"
	"pbrt/pkg/pbrt/film"
	"pbrt/pkg/pbrt/ray"
	"pbrt/pkg/pbrt/surface"
	"pbrt/pkg/pbrt/vec"
)

type RenderOptions struct {
	Width, Height   int
	SamplesPerPixel int
}

func Render(options RenderOptions, scene Scene, seed int64) film.Film {
	f := film.New(options.Width, options.Height)

	rng := rand.New(rand.NewSource(seed))

	for y := 0; y < options.Height; y++ {
		for x := 0; x < options.Width; x++ {
			sum := vec.Zero()
			for s := 0; s < options.SamplesPerPixel; s++ {
				u := (float64(x) + rng.Float64()) / float64(options.Width)
				v := (float64(y) + rng.Float64()) / float64(options.Height)

				r := scene.Camera.Ray(u, v, rng)

				color := pixelColor(r, scene.World)
				sum = sum.Add(color)
			}

			gammaCorrected := vec.New(
				math.Sqrt(sum.X/float64(options.SamplesPerPixel)),
				math.Sqrt(sum.Y/float64(options.SamplesPerPixel)),
				math.Sqrt(sum.Z/float64(options.SamplesPerPixel)),
			)
			f.Set(x, y, gammaCorrected)
		}
	}

	return f
}

func pixelColor(r ray.Ray, world surface.Surface) vec.Vec {
	if ok, isect := world.Intersect(r, math.SmallestNonzeroFloat32, math.Inf(1)); ok {
		return isect.Normal.Add(vec.One()).Scaled(0.5)
	}

	// background
	unitDirection := r.Direction.Normalized()
	t := (unitDirection.Y + 1) / 2
	return vec.New(0.82, 0.55, 0.24).Scaled(1 - t).Add(vec.New(0.46, 0.56, 0.66).Scaled(t))
}
