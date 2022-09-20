package pbrt

import (
	"math"
	"math/rand"
	"pbrt/pkg/pbrt/film"
	"pbrt/pkg/vec"
)

type RenderOptions struct {
	Width, Height   int
	SamplesPerPixel int
}

func Render(options RenderOptions) film.Film {
	f := film.New(options.Width, options.Height)

	for x := 0; x < options.Width; x++ {
		for y := 0; y < options.Height; y++ {
			sum := vec.Zero()
			for s := 0; s < options.SamplesPerPixel; s++ {
				i := (float64(x) + rand.Float64()) / float64(options.Width)
				j := (float64(y) + rand.Float64()) / float64(options.Height)

				color := pixelColor(i, j)
				sum = sum.Add(color)
			}

			gammaCorrected := vec.New(
				math.Sqrt(sum.X/float64(options.SamplesPerPixel)),
				math.Sqrt(sum.Y/float64(options.SamplesPerPixel)),
				math.Sqrt(sum.Z/float64(options.SamplesPerPixel)),
			)
			f.Set(x, y, gammaCorrected)

			// img.SetRGBA(x, y, color.RGBA{ir, ig, ib, 255})
		}
	}

	return f
}

func pixelColor(i, j float64) vec.Vec {
	return vec.New(i, j, 0.25)
}
