package main

import (
	"image/jpeg"
	"log"
	"os"
	"pbrt/pkg/pbrt"
	"pbrt/pkg/pbrt/camera"
	"pbrt/pkg/pbrt/film"
	"pbrt/pkg/pbrt/vec"
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
	aspectRatio := float64(options.Width) / float64(options.Height)

	lookFrom := vec.New(0, 1, -8)
	lookAt := vec.New(0, 1, 0)
	focusDist := lookAt.Sub(lookFrom).Len()

	scene := pbrt.Scene{
		Camera: camera.New(
			lookFrom, lookAt,
			vec.New(0, -1, 0),
			30,
			aspectRatio,
			0.01, focusDist,
		),
	}

	log.Println("starting render")

	t0 := time.Now()
	film := pbrt.Render(options, scene)
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
