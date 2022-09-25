package main

import (
	"image/jpeg"
	"log"
	"os"
	"time"

	"pbrt/cmd/scenes"
	"pbrt/pkg/pbrt"
	"pbrt/pkg/pbrt/film"
)

const (
	OUTPUT_FILENAME = "./output.jpg"
)

func start() {
	options := pbrt.RenderOptions{
		Width:           800,
		Height:          450,
		SamplesPerPixel: 8,
		MinDepth:        5,
	}
	aspectRatio := float64(options.Width) / float64(options.Height)

	log.Println("loading scene")
	scene := scenes.NewSpheres(aspectRatio)

	log.Println("starting render")

	rnd := pbrt.NewRenderer(options)
	t0 := time.Now()
	film := rnd.Render(scene, 0)
	d := time.Since(t0)

	log.Printf("finished render in %v", d)

	if OUTPUT_FILENAME != "" {
		writeFilm(film, OUTPUT_FILENAME)
		log.Printf("written film to file '%s'", OUTPUT_FILENAME)
	}
}

func writeFilm(film film.Film, filename string) {
	f, err := os.Create(OUTPUT_FILENAME)
	if err != nil {
		log.Fatalf("could not create output file '%s': %s", OUTPUT_FILENAME, err)
	}
	defer f.Close()

	img := film.ImageRGBA()

	if err := jpeg.Encode(f, img, &jpeg.Options{Quality: 100}); err != nil {
		log.Fatalf("could not encode output image: %s", err)
	}
}
