package main

import (
	"image"
	"image/jpeg"
	"log"
	"os"
	"pbrt/cmd/scenes"
	"pbrt/pkg/pbrt"
	"runtime"
	"time"
)

const (
	OUTPUT_FILENAME = "./output.jpg"
)

func start() {
	options := pbrt.RenderOptions{
		Width:           800,
		Height:          500,
		SamplesPerPixel: 2,
		MinDepth:        8,
	}
	aspectRatio := float64(options.Width) / float64(options.Height)

	log.Println("loading scene")
	scene := scenes.NewManySpheres(aspectRatio)

	rnd := pbrt.NewMultipass(options)
	numPasses := 4 * runtime.NumCPU()

	log.Println("starting render")
	t0 := time.Now()
	passes := rnd.Render(scene, numPasses)

	// consume passes and dump merged result to file
	n := 0
	for pass := range passes {
		n++
		if n%runtime.NumCPU() == 0 {
			if OUTPUT_FILENAME != "" {
				p := 100 * float64(n) / float64(numPasses)
				log.Printf("finished pass %d (%.1f%%), writing film to file '%s'", n, p, OUTPUT_FILENAME)
				img := pass.ImageRGBA()
				writeImage(img, OUTPUT_FILENAME)
			}
		}
	}
	d := time.Since(t0)

	log.Printf("finished %d passes in %v", numPasses, d)
	log.Printf("%v / pass", d/time.Duration(numPasses))
	stats := rnd.Stats()
	log.Printf("%s", stats)
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
