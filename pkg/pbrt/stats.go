package pbrt

import "fmt"

type RenderStats struct {
	Rays uint64
}

func (s RenderStats) String() string {
	return fmt.Sprintf("rays: %d", s.Rays)
}
