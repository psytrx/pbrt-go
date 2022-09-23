package pbrt

import (
	"pbrt/pkg/pbrt/camera"
	"pbrt/pkg/pbrt/surface"
)

type Scene struct {
	Camera camera.Camera
	World  surface.Surface
}
