package main

import (
	"image/jpeg"
	"log"
	"os"
	"pbrt/pkg/pbrt"
	"pbrt/pkg/pbrt/film"
	"time"
)

const (
	OUTPUT_FILENAME = "./output.jpg"
)

func start() {
	options := pbrt.RenderOptions{
		Width:           800,
		Height:          450,
		SamplesPerPixel: 32,
	}

	log.Println("starting render")

	t0 := time.Now()
	film := pbrt.Render(options)
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
