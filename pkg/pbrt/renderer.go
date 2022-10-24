package pbrt

import (
	"math"
	"math/rand"
	"pbrt/pkg/pbrt/vec"
	"sync/atomic"
)

type RenderOptions struct {
	Width, Height   int
	SamplesPerPixel int
	MinDepth        int
}

type Renderer struct {
	options RenderOptions
	rays    uint64
}

func NewRenderer(options RenderOptions) Renderer {
	return Renderer{options, 0}
}

func (rnd *Renderer) Render(scene Scene, seed int64) Film {
	f := NewFilm(rnd.options.Width, rnd.options.Height)

	rng := rand.New(rand.NewSource(seed))

	for y := 0; y < rnd.options.Height; y++ {
		for x := 0; x < rnd.options.Width; x++ {
			sum := vec.Zero()
			for s := 0; s < rnd.options.SamplesPerPixel; s++ {
				u := (float64(x) + rng.Float64()) / float64(rnd.options.Width)
				v := (float64(y) + rng.Float64()) / float64(rnd.options.Height)

				ray := scene.Camera.Ray(u, v, rng)
				atomic.AddUint64(&rnd.rays, 1)

				color := rnd.rayColor(&ray, scene, 0, rng)
				sum = sum.Add(color)
			}

			f.Set(x, y, sum.Scaled(1.0/float64(rnd.options.SamplesPerPixel)))
		}
	}

	return f
}

func (rnd Renderer) rayColor(r *Ray, scene Scene, depth int, rng *rand.Rand) vec.Vec {
	if depth > rnd.options.MinDepth {
		rrStopProp := 0.1
		if rng.Float64() < rrStopProp {
			return vec.Zero()
		}
	}

	if isected, isect := scene.World.Intersect(r, math.SmallestNonzeroFloat32, math.Inf(1)); isected {
		if ok, attenuation, scattered := isect.Material.Scatter(r, isect, rng); ok {
			rayColor := rnd.rayColor(scattered, scene, depth+1, rng)
			return attenuation.Mult(rayColor)
		}

		return vec.Zero()
	}

	return scene.Background.RayColor(r)
}

func (rnd Renderer) Stats() RenderStats {
	return RenderStats{
		Rays: rnd.rays,
	}
}
