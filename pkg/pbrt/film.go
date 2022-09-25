package pbrt

import (
	"image"
	"image/color"
	"math"

	"pbrt/pkg/pbrt/vec"
)

type Film struct {
	width, height int
	px            []vec.Vec
}

func NewFilm(width, height int) Film {
	px := make([]vec.Vec, width*height)
	return Film{width, height, px}
}

func (f Film) Set(x, y int, color vec.Vec) {
	idx := y*f.width + x
	f.px[idx] = color
}

func (f Film) Get(x, y int) vec.Vec {
	idx := y*f.width + x
	return f.px[idx]
}

func (f Film) ImageRGBA() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, f.width, f.height))
	for y := 0; y < f.height; y++ {
		for x := 0; x < f.width; x++ {
			px := f.Get(x, y)

			r := math.Sqrt(px.X)
			g := math.Sqrt(px.Y)
			b := math.Sqrt(px.Z)

			ir := uint8(255 * clamp(r, 0, 1))
			ig := uint8(255 * clamp(g, 0, 1))
			ib := uint8(255 * clamp(b, 0, 1))
			img.SetRGBA(x, y, color.RGBA{ir, ig, ib, 255})
		}
	}
	return img
}

func (f Film) Add(g Film, n int) Film {
	avg := NewFilm(f.width, f.height)
	for y := 0; y < f.height; y++ {
		for x := 0; x < f.width; x++ {
			fp := f.Get(x, y)
			gp := g.Get(x, y)

			sumP := fp.Add(gp)
			f.Set(x, y, sumP)
			avgP := sumP.Scaled(1.0 / float64(n))

			avg.Set(x, y, avgP)
		}
	}
	return avg
}

func clamp(v, min, max float64) float64 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}
