package main

import (
	"image/jpeg"
	"log"
	"os"
	"pbrt/pkg/pbrt"
	"pbrt/pkg/pbrt/film"
	"runtime"
	"runtime/pprof"
	"time"
)

const (
	CPU_PROFILE = true
	MEM_PROFILE = true
)

func main() {
	if CPU_PROFILE {
		cpuProf, err := os.Create("./cpu.prof")
		if err != nil {
			log.Fatalf("could not create CPU profile: %s", err)
		}
		defer cpuProf.Close()

		if err := pprof.StartCPUProfile(cpuProf); err != nil {
			log.Fatalf("could not start CPU profile: %s", err)
		}
		defer pprof.StopCPUProfile()
	}

	options := pbrt.RenderOptions{
		Width:           800,
		Height:          450,
		SamplesPerPixel: 32,
	}

	t0 := time.Now()
	film := pbrt.Render(options)
	d := time.Since(t0)

	log.Printf("render finished in %v", d)

	writeFilm(film)

	if MEM_PROFILE {
		memProf, err := os.Create("./mem.prof")
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer memProf.Close()

		runtime.GC()
		if err := pprof.WriteHeapProfile(memProf); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}

func writeFilm(film film.Film) {
	img := film.ImageRGBA()

	f, err := os.Create("./output.jpg")
	if err != nil {
		log.Fatalf("could not create output file: %s", err)
	}
	defer f.Close()

	if err := jpeg.Encode(f, img, &jpeg.Options{Quality: 100}); err != nil {
		log.Fatalf("could not encode output image: %s", err)
	}
}
