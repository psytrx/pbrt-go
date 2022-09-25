package pbrt

import (
	"runtime"
	"sync/atomic"
)

type MultipassRenderer struct {
	options RenderOptions
}

func NewMultipass(options RenderOptions) MultipassRenderer {
	return MultipassRenderer{options}
}

func (rnd MultipassRenderer) Render(scene Scene, numPasses int) chan Film {
	p := runtime.NumCPU()

	passes := make(chan Film, p)
	merged := make(chan Film, p)

	// launch workers
	seed := int64(0)
	for i := 0; i < p; i++ {
		go func() {
			for seed < int64(numPasses) {
				// fetch next seed
				currSeed := seed
				atomic.AddInt64(&seed, 1)

				r := NewRenderer(rnd.options)
				film := r.Render(scene, currSeed)

				passes <- film
			}
		}()
	}

	// merge passes
	go func() {
		sum := NewFilm(rnd.options.Width, rnd.options.Height)
		n := 0
		for pass := range passes {
			n++
			avg := sum.Add(pass, n)
			merged <- avg

			if n == numPasses {
				close(merged)
				return
			}
		}
	}()

	return merged
}
