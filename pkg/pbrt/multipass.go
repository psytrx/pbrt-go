package pbrt

import (
	"runtime"
	"sync/atomic"
)

type MultipassRenderer struct {
	options RenderOptions
	rays    uint64
}

func NewMultipass(options RenderOptions) MultipassRenderer {
	return MultipassRenderer{options, 0}
}

func (rnd *MultipassRenderer) Render(scene Scene, numPasses int) chan Film {
	nThreads := runtime.NumCPU()

	passes := make(chan Film, nThreads)
	merged := make(chan Film, nThreads)

	// launch workers
	seed := int64(0)
	for i := 0; i < nThreads; i++ {
		go func() {
			for seed < int64(numPasses) {
				// fetch next seed
				currSeed := seed
				atomic.AddInt64(&seed, 1)

				r := NewRenderer(rnd.options)
				film := r.Render(scene, currSeed)

				stats := r.Stats()
				atomic.AddUint64(&rnd.rays, stats.Rays)

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

func (rnd MultipassRenderer) Stats() RenderStats {
	return RenderStats{
		Rays: rnd.rays,
	}
}
