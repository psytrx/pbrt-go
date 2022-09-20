package vec

type Vec struct {
	X, Y, Z float64
}

func New(x, y, z float64) Vec {
	return Vec{x, y, z}
}

func (v Vec) Add(w Vec) Vec {
	return Vec{v.X + w.X, v.Y + w.Y, v.Z + w.Z}
}

func Zero() Vec {
	return Vec{0, 0, 0}
}
