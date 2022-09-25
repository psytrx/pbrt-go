package pbrt

type Scene struct {
	Camera     Camera
	World      Surface
	Background LinearGradientBackground
}
