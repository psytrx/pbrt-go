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
	MaxDepth        int
}

type Renderer struct {
	options RenderOptions
}

func NewRenderer(options RenderOptions) Renderer {
	return Renderer{options}
}

func (rnd Renderer) Render(scene Scene, seed int64) film.Film {
	f := film.New(rnd.options.Width, rnd.options.Height)

	rng := rand.New(rand.NewSource(seed))

	for y := 0; y < rnd.options.Height; y++ {
		for x := 0; x < rnd.options.Width; x++ {
			sum := vec.Zero()
			for s := 0; s < rnd.options.SamplesPerPixel; s++ {
				u := (float64(x) + rng.Float64()) / float64(rnd.options.Width)
				v := (float64(y) + rng.Float64()) / float64(rnd.options.Height)

				ray := scene.Camera.Ray(u, v, rng)

				color := rnd.rayColor(ray, scene.World, 0, rng)
				sum = sum.Add(color)
			}

			gammaCorrected := vec.New(
				math.Sqrt(sum.X/float64(rnd.options.SamplesPerPixel)),
				math.Sqrt(sum.Y/float64(rnd.options.SamplesPerPixel)),
				math.Sqrt(sum.Z/float64(rnd.options.SamplesPerPixel)),
			)
			f.Set(x, y, gammaCorrected)
		}
	}

	return f
}

func (rnd Renderer) rayColor(r ray.Ray, world surface.Surface, depth int, rng *rand.Rand) vec.Vec {
	if depth >= rnd.options.MaxDepth {
		return vec.Zero()
	}

	if ok, isect := world.Intersect(r, math.SmallestNonzeroFloat32, math.Inf(1)); ok {
		direction := isect.Normal.Add(vec.RandomInUnitSphere(rng))
		scattered := ray.New(isect.P, direction)
		return rnd.rayColor(scattered, world, depth-1, rng).Scaled(0.5)
	}

	// background
	unitDirection := r.Direction.Normalized()
	t := (unitDirection.Y + 1) / 2
	return vec.New(0.82, 0.55, 0.24).Scaled(1 - t).Add(vec.New(0.46, 0.56, 0.66).Scaled(t))
}
