package film

import (
	"image"
	"image/color"
	"pbrt/pkg/vec"
)

type Film struct {
	width, height int
	px            []vec.Vec
}

func New(width, height int) Film {
	px := make([]vec.Vec, width*height)
	return Film{width, height, px}
}

func (f Film) Set(x, y int, color vec.Vec) {
	idx := y*f.width + x
	f.px[idx] = color
}

func (f Film) ImageRGBA() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, f.width, f.height))
	for x := 0; x < f.width; x++ {
		for y := 0; y < f.height; y++ {
			idx := y*f.width + x
			px := f.px[idx]
			ir := uint8(255 * clamp(px.X, 0, 1))
			ig := uint8(255 * clamp(px.Y, 0, 1))
			ib := uint8(255 * clamp(px.Z, 0, 1))
			img.SetRGBA(x, y, color.RGBA{ir, ig, ib, 255})
		}
	}
	return img
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
