package main

import (
	"image"
	"image/jpeg"
	"log"
	"os"
	"time"

	"pbrt/cmd/scenes"
	"pbrt/pkg/pbrt"
)

const (
	OUTPUT_FILENAME = "./output.jpg"
)

func start() {
	options := pbrt.RenderOptions{
		Width:           800,
		Height:          450,
		SamplesPerPixel: 8,
		MinDepth:        8,
	}
	aspectRatio := float64(options.Width) / float64(options.Height)

	log.Println("loading scene")
	scene := scenes.NewSpheres(aspectRatio)

	rnd := pbrt.NewRenderer(options)

	log.Println("starting render")
	t0 := time.Now()
	film := rnd.Render(scene, 0)
	d := time.Since(t0)

	log.Printf("finished render in %v", d)

	if OUTPUT_FILENAME != "" {
		log.Printf("writing film to file '%s'", OUTPUT_FILENAME)
		img := film.ImageRGBA(options.SamplesPerPixel)
		writeImage(img, OUTPUT_FILENAME)
	}
}

func writeImage(img *image.RGBA, filename string) {
	f, err := os.Create(OUTPUT_FILENAME)
	if err != nil {
		log.Fatalf("could not create output file '%s': %s", OUTPUT_FILENAME, err)
	}
	defer f.Close()

	if err := jpeg.Encode(f, img, &jpeg.Options{Quality: 100}); err != nil {
		log.Fatalf("could not encode output image: %s", err)
	}
}
