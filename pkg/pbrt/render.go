package pbrt

import (
	"math"
	"math/rand"

	"pbrt/pkg/pbrt/film"
	"pbrt/pkg/pbrt/ray"
	"pbrt/pkg/pbrt/vec"
)

type RenderOptions struct {
	Width, Height   int
	SamplesPerPixel int
	MinDepth        int
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

				color := rnd.rayColor(ray, scene, 0, rng)
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

func (rnd Renderer) rayColor(r ray.Ray, scene Scene, depth int, rng *rand.Rand) vec.Vec {
	rrFactor := 1.0
	if depth >= rnd.options.MinDepth {
		rrStopProp := 0.1
		if rng.Float64() <= rrStopProp {
			return vec.Zero()
		}
		rrFactor = 1.0 / (1.0 - rrStopProp)
	}

	if ok, isect := scene.World.Intersect(r, math.SmallestNonzeroFloat32, math.Inf(1)); ok {
		direction := isect.Normal.Add(vec.RandomInUnitSphere(rng))
		scattered := ray.New(isect.P, direction)
		return rnd.rayColor(scattered, scene, depth+1, rng).Scaled(rrFactor * 0.5)
	}

	return scene.Background.RayColor(r)
}
