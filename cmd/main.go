package main

import (
	"image/jpeg"
	"log"
	"os"
	"pbrt/pkg/pbrt"
	"time"
)

func main() {
	options := pbrt.RenderOptions{
		Width:           800,
		Height:          450,
		SamplesPerPixel: 32,
	}

	t0 := time.Now()
	film := pbrt.Render(options)
	d := time.Since(t0)

	log.Printf("render finished in %v", d)

	img := film.ImageRGBA()

	f, err := os.Create("./output.jpg")
	if err != nil {
		log.Fatalf("could not create output file: %s", err)
	}
	if err := jpeg.Encode(f, img, &jpeg.Options{Quality: 100}); err != nil {
		log.Fatalf("could not encode output image: %s", err)
	}
}
