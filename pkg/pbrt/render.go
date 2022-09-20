package pbrt

import (
	"math"
	"math/rand"
	cam "pbrt/pkg/pbrt/camera"
	"pbrt/pkg/pbrt/film"
	"pbrt/pkg/pbrt/ray"
	"pbrt/pkg/pbrt/surface"
	"pbrt/pkg/pbrt/vec"
)

type RenderOptions struct {
	Width, Height   int
	SamplesPerPixel int
}

func Render(options RenderOptions) film.Film {
	f := film.New(options.Width, options.Height)

	aspectRatio := float64(options.Width) / float64(options.Height)
	c := cam.New(
		vec.New(0, 1, -8),
		vec.New(0, 1, 0),
		vec.New(0, -1, 0),
		30,
		aspectRatio,
		0, 4,
	)
	rng := rand.New(rand.NewSource(42))

	for y := 0; y < options.Height; y++ {
		for x := 0; x < options.Width; x++ {
			sum := vec.Zero()
			for s := 0; s < options.SamplesPerPixel; s++ {
				u := (float64(x) + rng.Float64()) / float64(options.Width)
				v := (float64(y) + rng.Float64()) / float64(options.Height)

				r := c.Ray(u, v)

				color := pixelColor(r)
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

func pixelColor(r ray.Ray) vec.Vec {
	s := surface.NewSphere(vec.New(0, 1, 0), 1)
	if t := s.Intersect(r, math.SmallestNonzeroFloat32, math.Inf(1)); t != nil {
		return vec.New(1, 0, 1)
	}

	unitDirection := r.Direction.Normalized()
	t := (unitDirection.Y + 1) / 2
	return vec.New(1, 1, 1).Scaled(1 - t).Add(vec.New(0.5, 0.7, 1).Scaled(t))
}
